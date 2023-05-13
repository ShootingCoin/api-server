package handler

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"github.com/ShootingCoin/api-server/core/common"
)

type WebSockerHandler struct {
	upgrader websocket.Upgrader
}

func (h *WebSockerHandler) ConnectWebSocket(c echo.Context) error {
	// Upgrade connection to WebSocket
	conn, err := h.upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	// Generate UUID for connection
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Errorln(err)
		return err
	}
	log.Infoln(fmt.Sprintf("New connection: %s", uuid.String()))

	// Store WebSocket connection and UUID
	common.StoreConnection(uuid.String(), conn)

	// Send UUID to frontend
	err = conn.WriteMessage(websocket.TextMessage, []byte(uuid.String()))
	if err != nil {
		log.Errorln(err)
		return err
	}

	// Continuously read messages from WebSocket
	go func() {
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err) {
					log.Infoln("Connection closed")
				} else {
					log.Errorln(err)
				}
				conn.Close()
				if common.IsConnected(uuid.String()) {
					common.DeleteConnection(uuid.String())
					log.Infoln(fmt.Sprintf("Deleted connection: %s", uuid.String()))
				}
				break
			}
		}
	}()

	// for debugging
	connections := common.ListConnections()
	for id, _ := range connections {
		log.Debugln(id)
	}

	return nil
}

func NewWebSocketHandler(upgrader websocket.Upgrader) *WebSockerHandler {
	return &WebSockerHandler{
		upgrader: upgrader,
	}
}

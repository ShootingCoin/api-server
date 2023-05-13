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

	// Store WebSocket connection and UUID
	err = common.StoreConnection(uuid.String(), conn)
	if err != nil {
		return err
	}
	log.Infoln(fmt.Sprintf("New connection: %s", uuid.String()))

	// Send UUID to frontend
	err = common.WriteMessage(uuid.String(), uuid.String())
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

				if err := common.DeleteConnection(uuid.String()); err != nil {
					log.Errorln(err)
					return
				}
				log.Infoln(fmt.Sprintf("Deleted connection: %s", uuid.String()))

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

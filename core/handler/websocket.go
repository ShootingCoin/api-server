package handler

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"github.com/ShootingCoin/api-server/core/common"
	"github.com/ShootingCoin/api-server/entity"
)

type WebSocketHandler struct {
	upgrader websocket.Upgrader
}

func (h *WebSocketHandler) ConnectWebSocket(c echo.Context) error {
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
	msgType, err := entity.ParseMessageType("wsId")
	if err != nil {
		return err
	}
	err = common.WriteMessage(uuid.String(), entity.Message{
		Type:    msgType.String(),
		Content: uuid.String(),
	})
	if err != nil {
		log.Errorln(err)
		return err
	}

	// Continuously read messages from WebSocket
	go func() {
		for {
			messageType, message, err := common.ReadMessage(uuid.String())
			if err != nil {
				if websocket.IsCloseError(err) {
					log.Infoln("Connection closed")
				} else {
					log.Errorln(err)
				}
				common.CloseConnection(uuid.String())

				// Delete connection and mark matched connection as unmatched
				if err := common.DeleteConnections(uuid.String()); err != nil {
					log.Errorln(err)
					return
				}
				log.Infoln(fmt.Sprintf("Deleted connection: %s", uuid.String()))

				break
			}

			// Check if connection is matched
			if matchUuid := common.GetMatched(uuid.String()); matchUuid != "" {
				// Start data transfer
				if messageType == websocket.TextMessage {
					// Forward the message to the matched client
					msgType, err := entity.ParseMessageType("gameInfo")
					if err != nil {
						log.Errorln(err)
					}
					err = common.WriteMessage(matchUuid, entity.Message{
						Type:    msgType.String(),
						Content: string(message),
					})
					if err != nil {
						log.Errorln("Failed to write message:", err)
					}
				}
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

func NewWebSocketHandler(upgrader websocket.Upgrader) *WebSocketHandler {
	return &WebSocketHandler{
		upgrader: upgrader,
	}
}

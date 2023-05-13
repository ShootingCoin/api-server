package common

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

// Map to store WebSocket connections and UUIDs
var connections = make(map[string]*websocket.Conn)
var connectionsMutex sync.Mutex
var games = make(map[string]string)
var gamesMutex sync.Mutex

func ListConnections() map[string]*websocket.Conn {
	return connections
}

func StoreConnection(uuid string, conn *websocket.Conn) error {
	if err := CheckConnection(uuid); err != nil {
		return err
	}

	connectionsMutex.Lock()
	defer connectionsMutex.Unlock()
	connections[uuid] = conn

	return nil
}

func DeleteConnection(uuid string) error {
	if err := CheckConnection(uuid); err != nil {
		return err
	}

	connectionsMutex.Lock()
	defer connectionsMutex.Unlock()
	delete(connections, uuid)

	return nil
}

func CheckConnection(uuid string) error {
	connectionsMutex.Lock()
	defer connectionsMutex.Unlock()

	if _, ok := connections[uuid]; !ok {
		return fmt.Errorf("connection not found: %s", uuid)
	}

	return nil
}

func WriteMessage(uuid string, msg string) error {
	if err := CheckConnection(uuid); err != nil {
		return err
	}

	if err := connections[uuid].WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
		return err
	}

	return nil
}

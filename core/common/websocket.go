package common

import (
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

func StoreConnection(uuid string, conn *websocket.Conn) {
	connectionsMutex.Lock()
	defer connectionsMutex.Unlock()
	connections[uuid] = conn
}

func DeleteConnection(uuid string) {
	connectionsMutex.Lock()
	defer connectionsMutex.Unlock()
	delete(connections, uuid)
}

func IsConnected(uuid string) bool {
	if _, ok := connections[uuid]; !ok {
		return false
	}

	return true
}

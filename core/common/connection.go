package common

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

// Map to store WebSocket connections and UUIDs
var connections = make(map[string]*websocket.Conn)
var connectionsMutex sync.Mutex
var matched = make(map[string]string)
var matchedMutex sync.Mutex

func ListConnections() map[string]*websocket.Conn {
	return connections
}

func StoreConnection(uuid string, conn *websocket.Conn) error {
	if err := checkConnection(uuid); err != nil {
		return err
	}

	connectionsMutex.Lock()
	defer connectionsMutex.Unlock()
	connections[uuid] = conn

	matchedMutex.Lock()
	defer matchedMutex.Unlock()
	matched[uuid] = ""

	return nil
}

func DeleteConnections(uuid string) error {
	if err := checkConnection(uuid); err != nil {
		return err
	}

	matchedMutex.Lock()
	defer matchedMutex.Unlock()

	// Delete matched connection
	if matchUuid := GetMatched(uuid); matchUuid != "" {
		delete(matched, matchUuid)
	}
	delete(matched, uuid)

	connectionsMutex.Lock()
	defer connectionsMutex.Unlock()
	delete(connections, uuid)

	return nil
}

func WriteMessage(uuid string, msg string) error {
	if err := checkConnection(uuid); err != nil {
		return err
	}

	if err := connections[uuid].WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
		return err
	}

	return nil
}

func ReadMessage(uuid string) (int, []byte, error) {
	if err := checkConnection(uuid); err != nil {
		return 0, nil, err
	}

	return connections[uuid].ReadMessage()
}

func CloseConnection(uuid string) error {
	if err := checkConnection(uuid); err != nil {
		return err
	}

	return connections[uuid].Close()
}

func GetMatched(uuid string) string {
	if err := checkMatched(uuid); err != nil {
		return ""
	}

	matchedMutex.Lock()
	defer matchedMutex.Unlock()

	return matched[uuid]
}

func SetMatched(reqUuid string, matchUuid string) error {
	if err := checkConnection(reqUuid); err != nil {
		return err
	}

	if err := checkConnection(matchUuid); err != nil {
		return err
	}

	matchedMutex.Lock()
	defer matchedMutex.Unlock()
	matched[reqUuid] = matchUuid
	matched[matchUuid] = reqUuid

	return nil
}

func checkConnection(uuid string) error {
	connectionsMutex.Lock()
	defer connectionsMutex.Unlock()

	if _, ok := connections[uuid]; !ok {
		return fmt.Errorf("connection not found: %s", uuid)
	}

	return nil
}

func checkMatched(uuid string) error {
	matchedMutex.Lock()
	defer matchedMutex.Unlock()

	if _, ok := matched[uuid]; !ok {
		return fmt.Errorf("connection not found: %s", uuid)
	}

	return nil
}

package entity

import "errors"

type MessageType string

const (
	GameIdMessageType      = MessageType("gameId")
	WebSocketIdMessageType = MessageType("wsId")
	GameInfoMessageType    = MessageType("gameInfo")
)

type Message struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

var AllMessageTypes = []MessageType{GameIdMessageType, WebSocketIdMessageType, GameInfoMessageType}

func (r MessageType) String() string {
	return string(r)
}

func ParseMessageType(msgType string) (MessageType, error) {
	for _, r := range AllMessageTypes {
		if msgType == r.String() {
			return r, nil
		}
	}
	return "", errors.New("cannot parse message type")
}

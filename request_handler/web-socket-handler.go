package request_handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type MessageWebSocketHandler struct {
	messageChannel chan string
}

type webSocketMessageWriter interface {
	WriteMessage(messageType int, message []byte) error
}

func NewMessageWebSocketHandler(messageChannel chan string) *MessageWebSocketHandler {
	return &MessageWebSocketHandler{messageChannel: messageChannel}
}

func (h *MessageWebSocketHandler) Handle(c *gin.Context) {
	if conn, err := wsUpgrader.Upgrade(c.Writer, c.Request, nil); err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v", err)
		return
	} else {
		h.messageSocketWriteHandler(conn)
	}
}

func (h *MessageWebSocketHandler) messageSocketWriteHandler(writer webSocketMessageWriter) {
	for message := range h.messageChannel {
		if err := writer.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			fmt.Printf("failed to write web socket message: %s", err)
		}
	}
}

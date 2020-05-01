package request_handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type MessageWebSocketHandler struct {
	messageChannel chan string
}

func NewMessageWebSocketHandler(messageChannel chan string) *MessageWebSocketHandler {
	return &MessageWebSocketHandler{messageChannel: messageChannel}
}

func (h *MessageWebSocketHandler) Handle(c *gin.Context) {
	h.wsHandler(c.Writer, c.Request, h.messageChannel)
}

func (h *MessageWebSocketHandler) wsHandler(w http.ResponseWriter, r *http.Request, messageChannel chan string) {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}
	for {
		for message := range messageChannel {
			if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
				fmt.Printf("failed to write web socket message: %s", err)
			}
		}
	}
}

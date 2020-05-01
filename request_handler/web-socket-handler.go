package request_handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type MessageWebSocketHandler struct {
}

func NewMessageWebSocketHandler() *MessageWebSocketHandler {
	return &MessageWebSocketHandler{}
}

func (h *MessageWebSocketHandler) Handle(c *gin.Context) {
	h.wsHandler(c.Writer, c.Request)
}

func (h *MessageWebSocketHandler) wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}
	for {
		time.Sleep(time.Second)
		if err := conn.WriteMessage(websocket.TextMessage, []byte(time.Now().Format("2020-12-31"))); err != nil {
			fmt.Printf("failed to write web socket message: %s", err)
		}
	}
}

package main

import (
	"github.com/erikfarhanmalik/simple-message/repository"
	request_handler "github.com/erikfarhanmalik/simple-message/request_handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	messageRepository := repository.NewInMemoryMessageRepository()
	messageChannel := make(chan string)
	messageRequestHandler := request_handler.NewMessageRequestHandler(messageRepository, messageChannel)
	router.POST("/messages", messageRequestHandler.SaveMessage)
	router.GET("/messages", messageRequestHandler.GetMessages)
	router.GET("/messages-board", messageRequestHandler.MessagesBoardPage)

	messageWebSocketHandler := request_handler.NewMessageWebSocketHandler(messageChannel)
	router.GET("/message-ws", messageWebSocketHandler.Handle)

	if err := router.Run(":3000"); err != nil {
		panic(err)
	}
}

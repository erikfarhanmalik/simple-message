package main

import (
	"github.com/erikfarhanmalik/simple-message/repository"
	request_handler "github.com/erikfarhanmalik/simple-message/request_handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	messageRequestHandler := request_handler.NewMessageRequestHandler(repository.NewInMemoryMessageRepository())
	router.POST("/messages", messageRequestHandler.SaveMessage)
	router.GET("/messages", messageRequestHandler.GetMessages)

	if err := router.Run(":3000"); err != nil {
		panic(err)
	}
}

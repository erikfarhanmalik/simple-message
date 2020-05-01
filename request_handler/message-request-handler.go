package request_handler

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/erikfarhanmalik/simple-message/dto"
	"github.com/erikfarhanmalik/simple-message/repository"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type MessageRequestHandler struct {
	messageRepository repository.MessageRepository
	messageChannel    chan<- string
}

func NewMessageRequestHandler(messageRepository repository.MessageRepository, messageChannel chan<- string) *MessageRequestHandler {
	return &MessageRequestHandler{
		messageRepository: messageRepository,
		messageChannel:    messageChannel,
	}
}

func (h *MessageRequestHandler) SaveMessage(c *gin.Context) {
	var message dto.Message
	if err := c.ShouldBindWith(&message, binding.JSON); err != nil {
		handleRequestError(c, http.StatusBadRequest, fmt.Errorf("invalid request payload: %s", err).Error())
		return
	}
	if err := h.messageRepository.Save(dto.CreateMessageModel(message)); err != nil {
		handleRequestError(c, http.StatusInternalServerError, fmt.Errorf("failed to save message: %s", err).Error())
		return
	}
	go func() { h.messageChannel <- message.Content }()
	c.JSON(200, gin.H{
		"message": message,
	})
}

func (h *MessageRequestHandler) GetMessages(c *gin.Context) {
	messages, err := h.messageRepository.GetAll()
	if err != nil {
		handleRequestError(c, http.StatusInternalServerError, fmt.Errorf("failed to get messages: %s", err).Error())
		return
	}
	c.JSON(200, gin.H{
		"message": dto.CreateMessagesFromModels(messages),
	})
}

func (h *MessageRequestHandler) MessagesBoardPage(c *gin.Context) {
	page, err := ioutil.ReadFile("./html_pages/message-board.html")
	if err != nil {
		handleRequestError(c, http.StatusInternalServerError, fmt.Errorf("could not open html page: %s", err).Error())
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", page)
}

func handleRequestError(c *gin.Context, errorCode int, message string) {
	c.AbortWithStatusJSON(errorCode, gin.H{
		"message": message,
	})
}

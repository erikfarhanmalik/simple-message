package request_handler

import (
	"fmt"
	"net/http"

	"github.com/erikfarhanmalik/simple-message/dto"
	"github.com/erikfarhanmalik/simple-message/repository"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type MessageRequestHandler struct {
	messageRepository repository.MessageRepository
}

func NewMessageRequestHandler(messageRepository repository.MessageRepository) MessageRequestHandler {
	return MessageRequestHandler{messageRepository: messageRepository}
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

func handleRequestError(c *gin.Context, errorCode int, message string) {
	c.AbortWithStatusJSON(errorCode, gin.H{
		"message": message,
	})
}

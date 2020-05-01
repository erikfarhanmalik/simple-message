package request_handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/erikfarhanmalik/simple-message/model"
	"github.com/erikfarhanmalik/simple-message/repository"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockMessageRepository struct {
}

func (r *mockMessageRepository) Save(message model.Message) error {
	return nil
}

func (r *mockMessageRepository) GetAll() ([]model.Message, error) {
	return []model.Message{}, errors.New("something went wrong")
}

func TestMessageRequestHandler_GetMessages(t *testing.T) {
	tests := []struct {
		name              string
		messages          []string
		messageRepository repository.MessageRepository
		wantResponse      string
		wantCode          int
	}{
		{
			"get messages should return proper http code and response on success",
			[]string{"message 1", "message 2", "message 3"},
			repository.NewInMemoryMessageRepository(),
			"{\"message\":[{\"content\":\"message 1\"},{\"content\":\"message 2\"},{\"content\":\"message 3\"}]}",
			200,
		},
		{
			"get messages should return proper http code and response on error",
			[]string{"message X", "message Y", "message Z"},
			&mockMessageRepository{},
			"{\"message\":\"failed to get messages: something went wrong\"}",
			500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			messageChannel := make(chan string)
			messageRequestHandler := NewMessageRequestHandler(tt.messageRepository, messageChannel)
			for _, message := range tt.messages {
				_ = tt.messageRepository.Save(model.Message{Content: message})
			}

			gin.SetMode(gin.TestMode)
			router := gin.New()
			router.GET("/messages", messageRequestHandler.GetMessages)

			response := httptest.NewRecorder()
			r, _ := http.NewRequest("GET", "/messages", nil)

			router.ServeHTTP(response, r)
			assert.Equal(t, response.Code, tt.wantCode)
			assert.Equal(t, response.Body.String(), tt.wantResponse)
		})
	}
}

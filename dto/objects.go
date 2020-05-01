package dto

import "github.com/erikfarhanmalik/simple-message/model"

type Message struct {
	Content string `json:"content" binding:"required"`
}

func CreateMessageModel(message Message) model.Message {
	return model.Message{Content: message.Content}
}

func CreateMessageFromModel(message model.Message) Message {
	return Message{Content: message.Content}
}

func CreateMessagesFromModels(messages []model.Message) []Message {
	result := make([]Message, 0)
	for _, message := range messages {
		result = append(result, CreateMessageFromModel(message))
	}
	return result
}

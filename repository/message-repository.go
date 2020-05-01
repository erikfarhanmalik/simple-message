package repository

import "github.com/erikfarhanmalik/simple-message/model"

type MessageRepository interface {
	Save(message model.Message) error
	GetAll() ([]model.Message, error)
}

type InMemoryMessageRepository struct {
	data []model.Message
}

func NewInMemoryMessageRepository() *InMemoryMessageRepository {
	return &InMemoryMessageRepository{data: []model.Message{}}
}

func (r *InMemoryMessageRepository) Save(message model.Message) error {
	r.data = append(r.data, message)
	return nil
}

func (r *InMemoryMessageRepository) GetAll() ([]model.Message, error) {
	return r.data, nil
}

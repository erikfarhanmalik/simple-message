package repository

import (
	"github.com/erikfarhanmalik/simple-message/model"
	"reflect"
	"testing"
)

func TestInMemoryMessageRepository_Save(t *testing.T) {
	tests := []struct {
		name         string
		message      model.Message
		addedMessage string
	}{
		{
			"save method should append the repository data",
			model.Message{Content: "some test message"},
			"some test message",
		},
		{
			"save method should append the repository data 2",
			model.Message{Content: "some other test message"},
			"some other test message",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewInMemoryMessageRepository()
			_ = r.Save(tt.message)
			if r.data[0].Content != tt.addedMessage {
				t.Errorf("Save() error, = %v is not added to the data", tt.addedMessage)
			}
		})
	}
}

func TestInMemoryMessageRepository_GetAll(t *testing.T) {
	tests := []struct {
		name string
		data []model.Message
		want []model.Message
	}{
		{
			"get all should retrieve all repository data",
			[]model.Message{
				{Content: "message 1"},
				{Content: "message 2"},
				{Content: "message 3"},
			},
			[]model.Message{
				{Content: "message 1"},
				{Content: "message 2"},
				{Content: "message 3"},
			},
		},
		{
			"get all should retrieve all repository data 2",
			[]model.Message{
				{Content: "message XXX"},
				{Content: "message YYY"},
				{Content: "message ZZZ"},
			},
			[]model.Message{
				{Content: "message XXX"},
				{Content: "message YYY"},
				{Content: "message ZZZ"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &InMemoryMessageRepository{data: tt.data}
			got, _ := r.GetAll()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

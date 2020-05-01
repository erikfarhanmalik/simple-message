package request_handler

import (
	"reflect"
	"sync"
	"testing"
)

type mockMebSocketMessageWriter struct {
	messages  []string
	waitGroup *sync.WaitGroup
}

func (w *mockMebSocketMessageWriter) WriteMessage(messageType int, message []byte) error {
	w.messages = append(w.messages, string(message))
	w.waitGroup.Done()
	return nil
}

func TestMessageWebSocketHandler_messageSocketWriteHandler(t *testing.T) {
	messageChannel := make(chan string)
	messageWebSocketHandler := NewMessageWebSocketHandler(messageChannel)
	var waitGroup sync.WaitGroup
	writer := &mockMebSocketMessageWriter{
		messages:  make([]string, 0),
		waitGroup: &waitGroup,
	}
	go messageWebSocketHandler.messageSocketWriteHandler(writer)

	inputMessage := []string{"message 1", "message 2", "message 3"}
	for _, message := range inputMessage {
		waitGroup.Add(1)
		messageChannel <- message
	}
	close(messageChannel)

	waitGroup.Wait()
	if !reflect.DeepEqual(writer.messages, inputMessage) {
		t.Errorf("messageSocketWriteHandler does not consume message through channel properly\n"+
			"input message: %+v, written message: %+v", inputMessage, writer.messages)
	}
}

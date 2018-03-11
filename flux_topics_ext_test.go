package flux_test

import (
	"testing"

	"github.com/dave/flux"

	"github.com/dave/ktest/assert"
)

type App struct {
	Dispatcher *flux.Dispatcher
	Messages   *MessageStore
	Topics     *TopicStore
}

func TestDispatcher(t *testing.T) {

	a := &App{}
	a.Messages = &MessageStore{app: a}
	a.Topics = &TopicStore{app: a}
	a.Dispatcher = flux.NewDispatcher(nil, a.Messages, a.Topics)

	done := a.Dispatcher.Dispatch(&AddMessage{Message: "a"})
	<-done
	done = a.Dispatcher.Dispatch(&AddTopic{Topic: "b", Message: "c"})
	<-done
	msg := a.Messages.GetMessages()
	assert.Equal(t, []string{"a", "c"}, msg)
}

type AddMessage struct {
	Message string
}

type MessageStore struct {
	app      *App
	messages []string
}

func (m *MessageStore) Handle(payload *flux.Payload) (finished bool) {
	switch action := payload.Action.(type) {
	case *AddTopic:
		payload.Wait(m.app.Topics)
		m.messages = append(m.messages, action.Message)
	case *AddMessage:
		m.messages = append(m.messages, action.Message)
	}
	return true
}

func (m *MessageStore) GetMessages() []string {
	return m.messages
}

type AddTopic struct {
	Topic   string
	Message string
}

type TopicStore struct {
	app    *App
	topics []string
}

func (m *TopicStore) Handle(payload *flux.Payload) (finished bool) {
	switch action := payload.Action.(type) {
	case *AddTopic:
		m.topics = append(m.topics, action.Topic)
	}
	return true
}

func (m *TopicStore) GetTopics() []string {
	return m.topics
}

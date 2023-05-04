package handlers

import (
	"log"

	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

type MessageLoggingHandler struct {
}

var _ EventHandler = &MessageLoggingHandler{}

func (s *MessageLoggingHandler) EventType() slackevents.EventsAPIType {
	return slackevents.Message
}

func (s *MessageLoggingHandler) Handle(
	ev *slackevents.EventsAPIEvent,
	cli *socketmode.Client,
) {
	log.Printf("MessageLogging: %+v\n", ev)
}

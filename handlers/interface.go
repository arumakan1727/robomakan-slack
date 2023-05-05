package handlers

import (
	"github.com/slack-go/slack/slackevents"
)

type EventHandler interface {
	EventType() slackevents.EventsAPIType
	Handle(event *slackevents.EventsAPIEvent)
}

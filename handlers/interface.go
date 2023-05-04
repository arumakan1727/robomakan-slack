package handlers

import (
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

type EventHandler interface {
	EventType() slackevents.EventsAPIType
	Handle(ev *slackevents.EventsAPIEvent, cli *socketmode.Client)
}

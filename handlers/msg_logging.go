package handlers

import (
	"log"

	"github.com/arumakan1727/robomakan-slack/withlogging"
	"github.com/k0kubun/pp/v3"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

type MessageLoggingHandler struct {
	cli *socketmode.Client
}

var _ EventHandler = &MessageLoggingHandler{}

func NewMessageLoggingHandler(cli *socketmode.Client) *MessageLoggingHandler {
	return &MessageLoggingHandler{cli}
}

func (h *MessageLoggingHandler) EventType() slackevents.EventsAPIType {
	return slackevents.Message
}

func (h *MessageLoggingHandler) Handle(event *slackevents.EventsAPIEvent) {
	ev, ok := withlogging.CastEvent[*slackevents.MessageEvent](event)
	if !ok {
		return
	}

	log.Println("------------------ MessageEvent -----------------")
	ch, _ := withlogging.GetChannelInfo(ev.Channel, h.cli)
	user, _ := withlogging.GetUserInfo(ev.Channel, h.cli)
	pp.Println(ev, ch, user)
}

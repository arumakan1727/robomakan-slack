package handlers

import (
	"log"
	"strings"

	"github.com/arumakan1727/robomakan-slack/config"
	"github.com/arumakan1727/robomakan-slack/util"
	"github.com/arumakan1727/robomakan-slack/withlogging"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

type TimesAllHandler struct {
	cfg *config.Config
	cli *socketmode.Client
}

var _ EventHandler = &TimesAllHandler{}

func NewTimesAllHandler(cfg *config.Config, cli *socketmode.Client) *TimesAllHandler {
	return &TimesAllHandler{
		cfg,
		cli,
	}
}

func (h *TimesAllHandler) EventType() slackevents.EventsAPIType {
	return slackevents.Message
}

func (h *TimesAllHandler) Handle(event *slackevents.EventsAPIEvent) {
	ev, ok := withlogging.CastEvent[*slackevents.MessageEvent](event)
	if !ok {
		return
	}

	shouldIgnore, _, user := h.shouldIgnore(ev)
	if shouldIgnore {
		log.Printf("TimesAllHandler: shouldIgnore=true")
		return
	}

	opts := []slack.MsgOption{
		slack.MsgOptionText(ev.Text, false),
		slack.MsgOptionUsername(user.Profile.DisplayName),
		slack.MsgOptionIconURL(user.Profile.Image192),
	}
	_, _, err := h.cli.PostMessage(h.cfg.TimesAllChannelID, opts...)
	if err != nil {
		log.Printf("TimesAllHandler: Failed to PostMessage: %+v\n", err)
	}
}

func (h *TimesAllHandler) shouldIgnore(ev *slackevents.MessageEvent) (bool, *slack.Channel, *slack.User) {
	if ev.User == "" {
		return true, nil, nil
	}

	user, err := withlogging.GetUserInfo(ev.User, h.cli)
	if err != nil {
		return true, nil, nil
	}
	if user.IsBot {
		return true, nil, nil
	}

	ch, err := withlogging.GetChannelInfo(ev.Channel, h.cli)
	if err != nil {
		return true, nil, nil
	}
	if !strings.HasPrefix(ch.Name, h.cfg.TimesChannelPrefix) {
		return true, nil, nil
	}

	// 「チャンネルにも投稿する」が指定されていないスレッドは無視
	if util.IsReplyMessage(ev) && ev.SubType != "thread_broadcast" {
		return true, nil, nil
	}

	return false, ch, user
}

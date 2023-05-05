package util

import "github.com/slack-go/slack/slackevents"

func IsReplyMessage(ev *slackevents.MessageEvent) bool {
	return len(ev.ThreadTimeStamp) > 0
}

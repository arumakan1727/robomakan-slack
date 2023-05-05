package withlogging

import (
	"log"
	"reflect"

	"github.com/arumakan1727/robomakan-slack/util"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func CastEvent[T any](e *slackevents.EventsAPIEvent) (T, bool) {
	x, ok := e.InnerEvent.Data.(T)
	if !ok {
		var t T
		log.Printf("[WARN] invalid cast EventsAPIEvent into %s: %+v\n",
			reflect.TypeOf(t).Name(),
			e,
		)
	}
	return x, ok
}

type conversationInfoGetter interface {
	GetConversationInfo(*slack.GetConversationInfoInput) (*slack.Channel, error)
}

func GetChannelInfo(channelID string, cli conversationInfoGetter) (*slack.Channel, error) {
	ch, err := cli.GetConversationInfo(&slack.GetConversationInfoInput{
		ChannelID: channelID,
	})
	if err != nil {
		log.Printf("[WARN] %s: Failed to GetConversationInfo: %v", util.GetCallerInfoStr(), err)
		return nil, err
	}
	return ch, nil
}

type userInfoGetter interface {
	GetUserInfo(user string) (*slack.User, error)
}

func GetUserInfo(userID string, cli userInfoGetter) (*slack.User, error) {
	user, err := cli.GetUserInfo(userID)
	if err != nil {
		log.Printf("[WARN] %s: Failed to GetUserInfo: %v", util.GetCallerInfoStr(), err)
		return nil, err
	}
	return user, nil
}

type permalinkGetter interface {
	GetPermalink(*slack.PermalinkParameters) (string, error)
}

func GetMsgPermalink(channelID, timestamp string, cli permalinkGetter) (string, error) {
	link, err := cli.GetPermalink(&slack.PermalinkParameters{
		Channel: channelID,
		Ts:      timestamp,
	})
	if err != nil {
		log.Printf("[WARN] %s: Failed to GetPermalink: %v", util.GetCallerInfoStr(), err)
		return "", err
	}
	return link, nil
}

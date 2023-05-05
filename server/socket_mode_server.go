package server

import (
	"context"
	"log"
	"os"

	"github.com/arumakan1727/robomakan-slack/config"
	"github.com/arumakan1727/robomakan-slack/handlers"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

type SocketModeServer struct {
	cli    *socketmode.Client
	hander *socketmode.SocketmodeHandler
}

func NewSocketModeServer(cfg *config.Config) (SocketModeServer, error) {
	if err := cfg.Validate(); err != nil {
		return SocketModeServer{}, err
	}

	debugEnabled := cfg.RunMode == config.ModeDebug

	baseClient := slack.New(
		cfg.SlackBotUserOAuthToken,
		slack.OptionDebug(debugEnabled),
		slack.OptionAppLevelToken(cfg.SlackAppLevelToken),
		slack.OptionLog(log.New(
			os.Stdout,
			"SlackClient: ",
			log.Lshortfile|log.LstdFlags,
		)),
	)

	sockModeClient := socketmode.New(baseClient)

	s := SocketModeServer{
		sockModeClient,
		socketmode.NewSocketmodeHandler(sockModeClient),
	}

	s.hander.HandleDefault(func(ev *socketmode.Event, cli *socketmode.Client) {
		log.Printf("[WARN] Unhandled event: %+v\n", ev)

		if ev.Type != socketmode.EventTypeHello && ev.Request != nil {
			cli.Ack(*ev.Request)
		}
	})

	return s, nil
}

func (s *SocketModeServer) RegisterEventHandlers(hs ...handlers.EventHandler) {
	for _, h := range hs {
		h := h
		s.hander.HandleEvents(h.EventType(), func(ev *socketmode.Event, cli *socketmode.Client) {
			evApiEv, ok := ev.Data.(slackevents.EventsAPIEvent)
			if !ok {
				log.Printf(
					"[WARN] (handler for %s) invalid cast into EventsAPIEvent: %v",
					h.EventType(), ev.Data,
				)
				return
			}

			cli.Ack(*ev.Request)
			h.Handle(&evApiEv)
		})
	}
}

func (s *SocketModeServer) SocketClient() *socketmode.Client {
	return s.cli
}

func (s *SocketModeServer) Serve(ctx context.Context) error {
	return s.hander.RunEventLoopContext(ctx)
}

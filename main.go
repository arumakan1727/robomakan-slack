package main

import (
	"context"
	"log"

	"github.com/arumakan1727/robomakan-slack/config"
	"github.com/arumakan1727/robomakan-slack/handlers"
	"github.com/arumakan1727/robomakan-slack/server"
)

func main() {
	cfg, err := config.LoadFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	s, err := server.NewSocketModeServer(cfg)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	cli := s.SocketClient()
	s.RegisterEventHandlers(
		handlers.NewMessageLoggingHandler(cli),
		handlers.NewTimesAllHandler(cfg, cli),
	)

	log.Println("Start serving...")
	log.Fatal(s.Serve(context.Background()))
}

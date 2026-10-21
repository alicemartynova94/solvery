package main

import (
	_ "github.com/lib/pq"
	"log"
	"solvery/07_lesson/chat_messenger/internal/api"
)

func main() {
	cfg, err := api.LoadConfig()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	api.ParseLimitsFlags(cfg)

	api.Run(cfg)
}

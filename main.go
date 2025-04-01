package main

import (
	"log"
	"os"

	tgClient "link-reminder-bot/clients/telegram"
	"link-reminder-bot/consumer/event_consumer"
	"link-reminder-bot/events/telegram"
	"link-reminder-bot/storage/files"

	"github.com/joho/godotenv"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := os.Getenv("TG_BOT_TOKEN")

	if token == "" {
		log.Fatal("token is not specified")
	}

	return token
}

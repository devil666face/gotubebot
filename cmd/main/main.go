package main

import (
	"log"

	"github.com/Devil666face/gotubebot/internal/bot"
	"github.com/Devil666face/gotubebot/internal/updater"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	bot, err := bot.Get()
	if err != nil {
		log.Fatal(err)
	}
	go updater.Start()
	bot.Start()
}

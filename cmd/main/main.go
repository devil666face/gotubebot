package main

import (
	"log"

	"github.com/Devil666face/gotubebot/internal/bot"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	bot, err := bot.StartBot()
	if err != nil {
		log.Fatal(err)
	}
	bot.Start()
}

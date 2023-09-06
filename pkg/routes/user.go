package routes

import (

	// "github.com/vitaliy-ukiru/fsm-telebot"
	"github.com/Devil666face/gotubebot/pkg/callbacks"
	"github.com/Devil666face/gotubebot/pkg/handlers"
	"gopkg.in/telebot.v3"
)

func setCommands(bot *telebot.Bot) {
	bot.Handle(callbacks.StartCommand, handlers.StartCommand)
}

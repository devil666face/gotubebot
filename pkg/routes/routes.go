package routes

import (
	"github.com/Devil666face/gotubebot/pkg/config"

	"github.com/vitaliy-ukiru/fsm-telebot"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

func New(bot *telebot.Bot, _ *fsm.Manager) {
	setMiddlewares(bot)
	setCommands(bot)
}

func setMiddlewares(bot *telebot.Bot) {
	if config.Cfg.Log {
		bot.Use(middleware.Logger())
	}
	bot.Use(middleware.AutoRespond())
}

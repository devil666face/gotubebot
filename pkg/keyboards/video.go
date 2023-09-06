package keyboards

import (
	"github.com/Devil666face/gotubebot/pkg/messages"
	"gopkg.in/telebot.v3"
)

var (
	CreateVideoBtn = telebot.ReplyButton{
		Text: messages.CreateVideo,
	}
	VideoMenu = &telebot.ReplyMarkup{
		ReplyKeyboard: [][]telebot.ReplyButton{
			{CreateVideoBtn},
			{BackBtn},
		},
		ResizeKeyboard: true,
	}
)

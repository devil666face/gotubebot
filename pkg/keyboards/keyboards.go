package keyboards

import (
	"github.com/Devil666face/gotubebot/pkg/messages"
	"gopkg.in/telebot.v3"
)

var (
	VideosBtn = telebot.ReplyButton{
		Text: messages.Videos,
	}
	PlaylistsBtn = telebot.ReplyButton{
		Text: messages.Playlists,
	}
	BackBtn = telebot.ReplyButton{
		Text: messages.Back,
	}
	MainMenu = &telebot.ReplyMarkup{
		ReplyKeyboard: [][]telebot.ReplyButton{
			{VideosBtn},
			{PlaylistsBtn},
			{BackBtn},
		},
		ResizeKeyboard: true,
	}
)

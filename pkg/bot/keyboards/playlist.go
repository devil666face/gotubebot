package keyboards

import (
	"fmt"

	"github.com/Devil666face/gotubebot/pkg/bot/callbacks"
	"github.com/Devil666face/gotubebot/pkg/bot/messages"
	"github.com/Devil666face/gotubebot/pkg/models"

	"gopkg.in/telebot.v3"
)

var (
	CreatePlaylistBtn = telebot.ReplyButton{
		Text: messages.CreatePlaylist,
	}
	PlaylistMenu = &telebot.ReplyMarkup{
		ReplyKeyboard: [][]telebot.ReplyButton{
			{CreatePlaylistBtn},
			{BackBtn},
		},
		ResizeKeyboard: true,
	}
)

func PlaylistsInline(playlists []models.Playlist) *telebot.ReplyMarkup {
	btns := [][]telebot.InlineButton{}
	for _, p := range playlists {
		btn := telebot.InlineButton{
			Text:   p.Title,
			Unique: fmt.Sprintf("%s:%d", callbacks.EditPlaylist, p.ID),
		}
		btns = append(btns, []telebot.InlineButton{btn})
	}
	return &telebot.ReplyMarkup{
		InlineKeyboard: btns,
		ResizeKeyboard: true,
	}
}

func EditPlaylistInline(id uint) *telebot.ReplyMarkup {
	showBtn := telebot.InlineButton{
		Text:   messages.ShowPlaylist,
		Unique: fmt.Sprintf("%s:%d", callbacks.ShowPlaylist, id),
	}
	scriptBtn := telebot.InlineButton{
		Text:   messages.GenScriptPlaylist,
		Unique: fmt.Sprintf("%s:%d", callbacks.GenScriptPlaylist, id),
	}
	deleteBtn := telebot.InlineButton{
		Text:   messages.Delete,
		Unique: fmt.Sprintf("%s:%d", callbacks.DeletePlaylist, id),
	}
	return &telebot.ReplyMarkup{
		InlineKeyboard: [][]telebot.InlineButton{
			{showBtn},
			{scriptBtn},
			{deleteBtn},
		},
		ResizeKeyboard: true,
	}

}

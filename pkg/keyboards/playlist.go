package keyboards

import (
	// "fmt"

	// "github.com/Devil666face/gotubebot/pkg/callbacks"
	"fmt"

	"github.com/Devil666face/gotubebot/pkg/callbacks"
	"github.com/Devil666face/gotubebot/pkg/messages"
	"github.com/Devil666face/gotubebot/pkg/models"

	// "github.com/Devil666face/gotubebot/pkg/models"
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
	updateBtn := telebot.InlineButton{
		Text:   messages.UpdatePlaylist,
		Unique: fmt.Sprintf("%s:%d", callbacks.UpdatePlaylist, id),
	}
	deleteBtn := telebot.InlineButton{
		Text:   messages.Delete,
		Unique: fmt.Sprintf("%s:%d", callbacks.DeletePlaylist, id),
	}
	return &telebot.ReplyMarkup{
		InlineKeyboard: [][]telebot.InlineButton{
			{showBtn},
			{updateBtn},
			{deleteBtn},
		},
		ResizeKeyboard: true,
	}

}

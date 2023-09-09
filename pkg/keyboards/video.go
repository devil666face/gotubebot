package keyboards

import (
	"fmt"

	"github.com/Devil666face/gotubebot/pkg/callbacks"
	"github.com/Devil666face/gotubebot/pkg/messages"
	"github.com/Devil666face/gotubebot/pkg/models"
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

func VideoListInline(videos []models.Video) *telebot.ReplyMarkup {
	btns := [][]telebot.InlineButton{}
	for _, v := range videos {
		btn := telebot.InlineButton{
			Text:   v.Title,
			Unique: fmt.Sprintf("%s:%d", callbacks.EditVideo, v.ID),
		}
		btns = append(btns, []telebot.InlineButton{btn})
	}
	return &telebot.ReplyMarkup{
		InlineKeyboard: btns,
		ResizeKeyboard: true,
	}
}

func UpdateOrDeleteVideoInline(id uint) *telebot.ReplyMarkup {
	update := telebot.InlineButton{
		Text:   messages.Update,
		Unique: fmt.Sprintf("%s:%d", callbacks.UpdateVideo, id),
	}
	delete := telebot.InlineButton{
		Text:   messages.Delete,
		Unique: fmt.Sprintf("%s:%d", callbacks.DeleteVideo, id),
	}
	return &telebot.ReplyMarkup{
		InlineKeyboard: [][]telebot.InlineButton{
			{update, delete},
		},
		ResizeKeyboard: true,
	}
}

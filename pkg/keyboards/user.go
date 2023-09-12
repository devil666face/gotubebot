package keyboards

import (
	"fmt"

	"github.com/Devil666face/gotubebot/pkg/callbacks"
	"github.com/Devil666face/gotubebot/pkg/messages"
	"gopkg.in/telebot.v3"
)

func InlineAddUser(id int64) *telebot.ReplyMarkup {
	confirmBtn := telebot.InlineButton{
		Text:   messages.ConfirmUser,
		Unique: fmt.Sprintf("%s:%d", callbacks.ConfirmUser, id),
	}
	ignoreBtn := telebot.InlineButton{
		Text:   messages.IgnoreUser,
		Unique: fmt.Sprintf("%s:%d", callbacks.IgnoreUser, id),
	}
	return &telebot.ReplyMarkup{
		InlineKeyboard: [][]telebot.InlineButton{
			{confirmBtn, ignoreBtn},
		},
		ResizeKeyboard: true,
	}
}

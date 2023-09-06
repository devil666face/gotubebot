package handlers

import (
	"log"

	"github.com/Devil666face/gotubebot/pkg/keyboards"
	"github.com/Devil666face/gotubebot/pkg/messages"
	"github.com/Devil666face/gotubebot/pkg/models"

	"gopkg.in/telebot.v3"
)

func deleteLastMessage(c telebot.Context) {
	if err := c.Delete(); err != nil {
		log.Print(err)
	}
}

func AskAdmins(c telebot.Context) {
	admins, err := models.GetAllAdmins()
	if err != nil {
		log.Print(messages.ErrGetAdminList)
	}
	for _, admin := range admins {
		if _, err := c.Bot().Send(
			&telebot.User{ID: int64(admin.TGID)},
			messages.AskAdminsForAddUser(c),
			keyboards.InlineAddUser(c.Chat().ID),
		); err != nil {
			log.Print(messages.ErrSendMessage(admin.Username))
		}
	}
}

func StartCommand(c telebot.Context) error {
	user := models.User{}
	if notfound := user.GetUserByTgID(c.Chat().ID); notfound != nil {
		user = models.User{
			TGID:     uint(c.Chat().ID),
			Username: c.Chat().Username,
		}
		if err := user.Create(); err != nil {
			return c.Send(messages.ErrCreateUser(c))
		}
	}
	if !user.IsAllow {
		AskAdmins(c)
		return c.Send(
			messages.SuccessfulCreateUser(c), telebot.RemoveKeyboard,
		)
	}
	return c.Send(
		messages.ErrUserAlreadyCreate(c), keyboards.MainMenu,
	)
}

package handlers

import (
	"github.com/Devil666face/gotubebot/pkg/callbacks"
	"github.com/Devil666face/gotubebot/pkg/keyboards"
	"github.com/Devil666face/gotubebot/pkg/messages"
	"github.com/Devil666face/gotubebot/pkg/models"
	"github.com/Devil666face/gotubebot/pkg/utils"

	"github.com/vitaliy-ukiru/fsm-telebot"
	telebot "gopkg.in/telebot.v3"
)

func OnConfirmUser(c telebot.Context, _ fsm.Context) error {
	defer delMes(c)
	id := utils.ToInt64(c.Get(callbacks.CallbackVal))
	user := models.User{}
	if notfound := user.GetUserByTgID(id); notfound != nil {
		return c.Send(messages.ErrUserNotFound(id))
	}
	user.IsAllow = true
	if err := user.Update(); err != nil {
		return c.Send(messages.ErrUserUpdate(user.Username))
	}
	if _, err := c.Bot().Send(
		&telebot.User{ID: int64(user.TGID)},
		messages.PermissionsForUserAdded(user.Username),
		keyboards.MainMenu); err != nil {
		return c.Send(messages.ErrSendMessage(user.Username))
	}
	return c.Send(
		messages.SuccessfulUpdateUser(user.Username), keyboards.MainMenu,
	)
}

func OnIgnoreUser(c telebot.Context, _ fsm.Context) error {
	return c.Delete()
}

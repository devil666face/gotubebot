package handlers

import (
	"log"

	"github.com/Devil666face/gotubebot/pkg/keyboard"
	"github.com/Devil666face/gotubebot/pkg/messages"
	"github.com/Devil666face/gotubebot/pkg/models"
	"github.com/Devil666face/gotubebot/pkg/utils"

	"github.com/vitaliy-ukiru/fsm-telebot"
	telebot "gopkg.in/telebot.v3"
)

func OnConfirmUser(c telebot.Context) error {
	id := utils.ToInt64(c.Get(CALLBACK_VAL))
	user := models.User{}
	if notfound := user.GetUserByTgID(id); notfound != nil {
		return c.Send(m.ErrUserNotFound(id))
	}
	user.IsAllow = true
	if err := user.Update(); err != nil {
		return c.Send(m.ErrUserUpdate(user.Username))
	}
	if err := c.Delete(); err != nil {
		log.Print(err)
	}
	if _, err := c.Bot().Send(&telebot.User{ID: int64(user.TGID)}, m.PermissionsForUserAdded(user.Username), kb.Menu); err != nil {
		return c.Send(m.ErrSendMessage(user.Username))
	}
	return c.Send(m.SuccessfulUpdateUser(user.Username), kb.Menu)
}

func OnIgnoreUser(c telebot.Context) error {
	return c.Delete()
}

func OnBackBtn(c telebot.Context, f fsm.Context) error {
	if err := f.Finish(true); err != nil {
		log.Print(err)
	}
	return c.Send(m.GO_BACK, kb.Menu)
}

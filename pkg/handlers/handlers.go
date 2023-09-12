package handlers

import (
	"log"

	"github.com/Devil666face/gotubebot/pkg/keyboards"
	"github.com/Devil666face/gotubebot/pkg/messages"
	"github.com/Devil666face/gotubebot/pkg/models"
	"github.com/vitaliy-ukiru/fsm-telebot"

	telebot "gopkg.in/telebot.v3"
)

func delMes(c telebot.Context) {
	if err := c.Delete(); err != nil {
		log.Print(err)
	}
}

func finish(s fsm.Context) {
	if err := s.Finish(true); err != nil {
		log.Print(err)
	}
}

func setState(s fsm.Context, state fsm.State) {
	if err := s.Set(state); err != nil {
		log.Print(err)
	}
}

func askAdmins(c telebot.Context) {
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

func OnStartCommand(c telebot.Context, _ fsm.Context) error {
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
		askAdmins(c)
		return c.Send(
			messages.SuccessfulCreateUser(c), telebot.RemoveKeyboard,
		)
	}
	return c.Send(
		messages.ErrUserAlreadyCreate(c), keyboards.MainMenu,
	)
}

func OnBackBtn(c telebot.Context, s fsm.Context) error {
	defer finish(s)
	return c.Send(
		messages.GoBack, keyboards.MainMenu,
	)
}

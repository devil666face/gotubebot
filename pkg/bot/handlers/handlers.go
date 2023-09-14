package handlers

import (
	"log"

	"github.com/Devil666face/gotubebot/pkg/bot/keyboards"
	"github.com/Devil666face/gotubebot/pkg/bot/messages"
	"github.com/Devil666face/gotubebot/pkg/models"
	"github.com/Devil666face/gotubebot/pkg/utils"
	"github.com/vitaliy-ukiru/fsm-telebot" //nolint:misspell

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

func OnText(c telebot.Context, s fsm.Context) error {
	if err := utils.ValidateYtURL(c.Message().Text); err != nil {
		return c.Send(messages.ErrParseYtURL)
	}
	defer finish(s)

	video := models.Video{
		URL: c.Message().Text,
	}

	if err := video.ParseYt(); err != nil {
		log.Print(err)
		return c.Send(messages.ErrLoadVideoFromYt, keyboards.VideoMenu)
	}
	return c.Send(video.String(), keyboards.VideoMenu)
}

func OnBackBtn(c telebot.Context, s fsm.Context) error {
	defer finish(s)
	return c.Send(
		messages.Back, keyboards.MainMenu,
	)
}

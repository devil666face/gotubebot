package handlers

import (
	"log"
	"strings"

	"github.com/Devil666face/gotubebot/pkg/callbacks"
	"github.com/Devil666face/gotubebot/pkg/keyboards"
	"github.com/Devil666face/gotubebot/pkg/messages"
	"github.com/Devil666face/gotubebot/pkg/models"
	"github.com/vitaliy-ukiru/fsm-telebot"

	telebot "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

func CallbackKeyValueMiddleware(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		if c.Callback() != nil {
			callbackData := strings.ReplaceAll(c.Callback().Data, "\f", "")
			callbackPayload := strings.Split(callbackData, ":")
			c.Set(callbacks.CallbackKey, callbackPayload[0])
			c.Set(callbacks.CallbackVal, callbackPayload[1])
		}
		return next(c)
	}
}

func AdminOnlyMiddleware(next telebot.HandlerFunc) telebot.HandlerFunc {
	return getPermissionHandleFunc(models.GetAllAdmins, next)
}

func AllowOnlyMiddleware(next telebot.HandlerFunc) telebot.HandlerFunc {
	return getPermissionHandleFunc(models.GetAllAllows, next)
}

func getPermissionHandleFunc(selectorFunc func() ([]models.User, error), next telebot.HandlerFunc) telebot.HandlerFunc {
	chats, err := models.GetChatIdsForSelector(selectorFunc)
	if err != nil {
		log.Print(err)
		return nil
	}
	return middleware.Whitelist(chats...)(next)

}

func AdminOnlyDecorator(next fsm.Handler) fsm.Handler {
	return getPermissonHanler(models.GetAllAdmins, next)
}

func AllowOnlyDecorator(next fsm.Handler) fsm.Handler {
	return getPermissonHanler(models.GetAllAllows, next)
}

func getPermissonHanler(selectorFunc func() ([]models.User, error), next fsm.Handler) fsm.Handler {
	return func(c telebot.Context, s fsm.Context) error {
		chats, err := models.GetChatIdsForSelector(selectorFunc)
		if err != nil {
			log.Print(err)
			return nil
		}
		for _, id := range chats {
			if c.Chat().ID == id {
				return next(c, s)
			}
		}
		return nil
	}
}

func UserInContextDecorator(next fsm.Handler) fsm.Handler {
	return func(c telebot.Context, s fsm.Context) error {
		user := models.User{}
		if err := user.GetUserByTgID(c.Chat().ID); err != nil {
			log.Print(err)
			return c.Send(messages.ErrGetUser, keyboards.MainMenu)
		}
		c.Set(callbacks.UserKey, user)
		return next(c, s)
	}
}

package handlers

import (
	"log"
	"strings"

	"github.com/Devil666face/gotubebot/pkg/callbacks"
	"github.com/Devil666face/gotubebot/pkg/models"

	telebot "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

func CallbackKeyValueMiddleware(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		if c.Callback() != nil {
			r := strings.ReplaceAll(c.Callback().Data, "\f", "")
			data := strings.Split(r, ":")
			c.Set(callbacks.CallbackKey, data[0])
			c.Set(callbacks.CallbackVal, data[1])
		}
		return next(c)
	}
}

func permission(selectorFunc func() ([]models.User, error), next telebot.HandlerFunc) telebot.HandlerFunc {
	chats, err := models.GetChatIdsForSelector(selectorFunc)
	if err != nil {
		log.Print(err)
		return middleware.Whitelist()(next)
	}
	return middleware.Whitelist(chats...)(next)

}

func AdminOnlyMiddleware(next telebot.HandlerFunc) telebot.HandlerFunc {
	return permission(models.GetAllAdmins, next)
}

func AllowOnlyMiddleware(next telebot.HandlerFunc) telebot.HandlerFunc {
	return permission(models.GetAllAllows, next)
}

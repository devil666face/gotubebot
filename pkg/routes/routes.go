package routes

import (
	"github.com/Devil666face/gotubebot/pkg/callbacks"
	"github.com/Devil666face/gotubebot/pkg/config"
	"github.com/Devil666face/gotubebot/pkg/handlers"
	"github.com/Devil666face/gotubebot/pkg/keyboards"

	"github.com/vitaliy-ukiru/fsm-telebot"
	telebot "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

var callbackMap = map[string]func(telebot.Context, fsm.Context) error{
	callbacks.ConfirmUser: handlers.AdminOnlyDecorator(handlers.OnConfirmUser),
	callbacks.IgnoreUser:  handlers.AdminOnlyDecorator(handlers.OnIgnoreUser),
}

type Manager struct {
	*fsm.Manager
}

func New(manager *Manager) {
	manager.setMiddelwares()
	manager.setFreeCommands()
	manager.setCallbacks()
	manager.setVideoRoutes()

	manager.Use(handlers.AllowOnlyMiddleware)
	manager.Bind(
		&keyboards.BackBtn,
		fsm.AnyState,
		handlers.OnBackBtn,
		handlers.AllowOnlyMiddleware,
	)
}

func (manager *Manager) setMiddelwares() {
	if config.Cfg.Log {
		manager.Use(middleware.Logger())
	}
	manager.Use(middleware.AutoRespond())
}

func (manager *Manager) setFreeCommands() {
	manager.Bind(callbacks.StartCommand, fsm.AnyState, handlers.OnStartCommand)
}

func (manager *Manager) setCallbacks() {
	manager.Bind(
		telebot.OnCallback,
		fsm.AnyState,
		func(c telebot.Context, s fsm.Context) error {
			callback := c.Get(callbacks.CallbackKey).(string)
			if f, ok := callbackMap[callback]; ok {
				return f(c, s)
			}
			return nil
		},
		handlers.CallbackKeyValueMiddleware,
		handlers.AllowOnlyMiddleware,
	)
}

func CallbackHandler(c telebot.Context, s fsm.Context) error {
	callback := c.Get(callbacks.CallbackKey).(string)
	if f, ok := callbackMap[callback]; ok {
		return f(c, s)
	}
	return nil
}

package routes

import (
	"github.com/Devil666face/gotubebot/pkg/bot/callbacks"
	"github.com/Devil666face/gotubebot/pkg/bot/handlers"
	"github.com/Devil666face/gotubebot/pkg/bot/keyboards"
	"github.com/Devil666face/gotubebot/pkg/config"

	"github.com/vitaliy-ukiru/fsm-telebot"
	telebot "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

var callbackMap = map[string]func(telebot.Context, fsm.Context) error{
	callbacks.ConfirmUser:       handlers.AdminOnlyDecorator(handlers.OnConfirmUser),
	callbacks.IgnoreUser:        handlers.AdminOnlyDecorator(handlers.OnIgnoreUser),
	callbacks.EditVideo:         handlers.AllowOnlyDecorator(handlers.OnEditVideoInlineBtn),
	callbacks.DeleteVideo:       handlers.AllowOnlyDecorator(handlers.UserInCtxDecorator(handlers.OnDeleteVideoInlineBtn)),
	callbacks.EditPlaylist:      handlers.AllowOnlyDecorator(handlers.OnEditPlaylistInlineBtn),
	callbacks.ShowPlaylist:      handlers.AllowOnlyDecorator(handlers.OnShowPlaylistInlineBtn),
	callbacks.GenScriptPlaylist: handlers.AllowOnlyDecorator(handlers.OnGenScriptPlaylistInlineBtn),
	callbacks.DeletePlaylist:    handlers.AllowOnlyDecorator(handlers.UserInCtxDecorator(handlers.OnDeletePlaylistInlineBtn)),
}

type Manager struct {
	*fsm.Manager
}

func New(manager *Manager) {
	manager.setMiddelwares()
	manager.setFreeCommands()
	manager.setCallbacks()

	manager.Use(handlers.AllowOnlyMiddleware)

	manager.setVideoRoutes()
	manager.setPlaylistRoutes()
	manager.Bind(
		&keyboards.BackBtn,
		fsm.AnyState,
		handlers.OnBackBtn,
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

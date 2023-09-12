package routes

import (
	"github.com/Devil666face/gotubebot/pkg/bot/handlers"
	"github.com/Devil666face/gotubebot/pkg/bot/keyboards"
	"github.com/vitaliy-ukiru/fsm-telebot"
	"gopkg.in/telebot.v3"
)

func (manager *Manager) setPlaylistRoutes() {
	manager.Bind(&keyboards.PlaylistsBtn, fsm.AnyState, handlers.UserInCtxDecorator(handlers.OnPlaylistsBtn))
	manager.Bind(&keyboards.CreatePlaylistBtn, fsm.AnyState, handlers.OnCreatePlaylistBtn)
	manager.Bind(telebot.OnText, handlers.CreatePlaylistState, handlers.UserInCtxDecorator(handlers.OnRecivePlaylistURL))
}

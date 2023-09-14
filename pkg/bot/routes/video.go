package routes

import (
	"github.com/Devil666face/gotubebot/pkg/bot/handlers"
	"github.com/Devil666face/gotubebot/pkg/bot/keyboards"
	"github.com/vitaliy-ukiru/fsm-telebot" //nolint:misspell
	"gopkg.in/telebot.v3"
)

func (manager *Manager) setVideoRoutes() {
	manager.Bind(&keyboards.VideosBtn, fsm.AnyState, handlers.UserInCtxDecorator(handlers.OnVideosBtn))
	manager.Bind(&keyboards.CreateVideoBtn, fsm.AnyState, handlers.OnCreateVideoBtn)
	manager.Bind(telebot.OnText, handlers.CreateVideoState, handlers.UserInCtxDecorator(handlers.OnReciveVideoURL))
}

package routes

import (
	"github.com/Devil666face/gotubebot/pkg/handlers"
	"github.com/Devil666face/gotubebot/pkg/keyboards"
	"github.com/vitaliy-ukiru/fsm-telebot"
	"gopkg.in/telebot.v3"
)

func (manager *Manager) setVideoRoutes() {
	manager.Bind(&keyboards.VideosBtn, fsm.AnyState, handlers.UserInContextDecorator(handlers.OnVideosBtn))
	manager.Bind(&keyboards.CreateVideoBtn, fsm.AnyState, handlers.OnCreateVideoBtn)
	manager.Bind(telebot.OnText, handlers.CreateVideoState, handlers.UserInContextDecorator(handlers.OnReciveVideoURL))
}

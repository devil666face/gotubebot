package handlers

import (
	"github.com/Devil666face/gotubebot/pkg/keyboards"
	"github.com/Devil666face/gotubebot/pkg/messages"
	"github.com/vitaliy-ukiru/fsm-telebot"
	telebot "gopkg.in/telebot.v3"
)

var (
	createVideoGroup = fsm.NewStateGroup("video")
	CreateVideoState = createVideoGroup.New("create")
)

func OnVideosBtn(c telebot.Context, _ fsm.Context) error {
	return c.Send(messages.ChangeVideo, keyboards.VideoMenu)
}

func OnCreateVideoBtn(c telebot.Context, s fsm.Context) error {
	defer setState(s, CreateVideoState)
	return c.Send(messages.SendVideoUrl, keyboards.BackMenu)
}

func OnReciveVideoUrl(c telebot.Context, s fsm.Context) error {
	defer finish(s)
	return c.Send(c.Message().Text, keyboards.VideoMenu)
}

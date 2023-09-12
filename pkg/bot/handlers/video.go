package handlers

import (
	"log"

	"github.com/Devil666face/gotubebot/pkg/bot/callbacks"
	"github.com/Devil666face/gotubebot/pkg/bot/keyboards"
	"github.com/Devil666face/gotubebot/pkg/bot/messages"
	"github.com/Devil666face/gotubebot/pkg/models"
	"github.com/Devil666face/gotubebot/pkg/utils"

	"github.com/vitaliy-ukiru/fsm-telebot"
	telebot "gopkg.in/telebot.v3"
)

var (
	createVideoGroup = fsm.NewStateGroup("video")
	CreateVideoState = createVideoGroup.New("create")
)

func OnVideosBtn(c telebot.Context, _ fsm.Context) error {
	user := c.Get(callbacks.UserKey).(models.User)

	videos, err := models.GetAllVideosForUser(user.ID)
	if err != nil {
		log.Print(err)
	}

	if err := c.Send(messages.ChangeVideo, keyboards.VideoMenu); err != nil {
		log.Print(err)
	}
	return c.Send(messages.VideoList, keyboards.VideoListInline(videos))
}

func OnCreateVideoBtn(c telebot.Context, s fsm.Context) error {
	defer setState(s, CreateVideoState)
	return c.Send(messages.SendVideoURL, keyboards.BackMenu)
}

func OnReciveVideoURL(c telebot.Context, s fsm.Context) error {
	if err := utils.ValidateYtURL(c.Message().Text); err != nil {
		return c.Send(messages.ErrParseYtURL)
	}
	defer finish(s)

	user := c.Get(callbacks.UserKey).(models.User)

	video := models.Video{
		URL:    c.Message().Text,
		UserID: user.ID,
	}

	if err := video.ParseYt(); err != nil {
		return c.Send(messages.ErrLoadVideoFromYt, keyboards.VideoMenu)
	}
	if err := video.Create(); err != nil {
		log.Print(err)
	}
	return c.Send(video.String(), keyboards.VideoMenu)
}

func OnEditVideoInlineBtn(c telebot.Context, _ fsm.Context) error {
	defer delMes(c)
	video := models.Video{}
	if err := video.Get(utils.ToUint(c.Get(callbacks.CallbackVal))); err != nil {
		log.Print(err)
		return c.Send(messages.ErrGetVideo, keyboards.MainMenu)
	}
	return c.Send(video.String(), keyboards.UpdateOrDeleteVideoInline(video.ID))
}

func OnUpdateVideoInlineBtn(c telebot.Context, _ fsm.Context) error {
	defer delMes(c)
	video := models.Video{}
	if err := video.Get(utils.ToUint(c.Get(callbacks.CallbackVal))); err != nil {
		return c.Send(messages.ErrGetVideo, keyboards.MainMenu)
	}
	if err := video.ParseYt(); err != nil {
		return c.Send(messages.ErrLoadVideoFromYt, keyboards.VideoMenu)
	}
	if err := video.Update(); err != nil {
		log.Print(err)
	}
	return inlineVideosForUser(c)
}

func OnDeleteVideoInlineBtn(c telebot.Context, _ fsm.Context) error {
	defer delMes(c)
	video := models.Video{}
	if err := video.Get(utils.ToUint(c.Get(callbacks.CallbackVal))); err != nil {
		return c.Send(messages.ErrGetVideo, keyboards.MainMenu)
	}
	if err := video.Delete(); err != nil {
		log.Print(err)
		return c.Send(messages.ErrDeleteVideo, keyboards.MainMenu)
	}
	return inlineVideosForUser(c)
}

func inlineVideosForUser(c telebot.Context) error {
	user := c.Get(callbacks.UserKey).(models.User)
	videos, err := models.GetAllVideosForUser(user.ID)
	if err != nil {
		log.Print(err)
	}
	return c.Send(messages.SuccessfulUpdateVideo, keyboards.VideoListInline(videos))
}

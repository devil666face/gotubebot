package handlers

import (
	// "fmt"

	"log"

	"github.com/Devil666face/gotubebot/pkg/callbacks"
	"github.com/Devil666face/gotubebot/pkg/keyboards"
	"github.com/Devil666face/gotubebot/pkg/messages"
	"github.com/Devil666face/gotubebot/pkg/models"
	"github.com/Devil666face/gotubebot/pkg/utils"

	"github.com/vitaliy-ukiru/fsm-telebot"
	telebot "gopkg.in/telebot.v3"
)

var (
	createPlaylistGroup = fsm.NewStateGroup("playlist")
	CreatePlaylistState = createPlaylistGroup.New("create")
)

func OnPlaylistsBtn(c telebot.Context, _ fsm.Context) error {
	user := c.Get(callbacks.UserKey).(models.User)

	playlists, err := models.GetAllPlaylistsForUser(user.ID)
	if err != nil {
		log.Print(err)
	}

	if err := c.Send(messages.PlaylistList, keyboards.PlaylistMenu); err != nil {
		log.Print(err)
	}
	return c.Send(messages.ChangePlaylist, keyboards.PlaylistsInline(playlists))
}

func OnCreatePlaylistBtn(c telebot.Context, s fsm.Context) error {
	defer setState(s, CreatePlaylistState)
	return c.Send(messages.SendPlaylistUrl, keyboards.BackMenu)
}

func OnRecivePlaylistUrl(c telebot.Context, s fsm.Context) error {
	if err := utils.ValidateYtURL(c.Message().Text); err != nil {
		return c.Send(messages.ErrParseYtURL)
	}
	defer finish(s)

	user := c.Get(callbacks.UserKey).(models.User)

	playlist := models.Playlist{
		Url:    c.Message().Text,
		UserID: user.ID,
	}

	videos, err := playlist.ParseYt()
	if err != nil {
		log.Print(err)
		return c.Send(messages.ErrLoadPlaylistFromYt, keyboards.PlaylistMenu)
	}

	if err := playlist.Create(); err != nil {
		log.Print(err)
		return c.Send(messages.ErrLoadPlaylistFromYt, keyboards.PlaylistMenu)
	}

	for _, video := range videos {
		go func(c telebot.Context, video models.Video) error {
			if err := video.ParseYt(); err != nil {
				log.Print(err)
				return c.Send(messages.ErrParseYtURL)
			}
			video.PlaylistID = playlist.ID
			if err := video.Create(); err != nil {
				log.Print(err)
				return c.Send(messages.ErrLoadVideoFromYt)
			}
			return c.Send(video.String())
		}(c, video)
	}

	return c.Send(messages.SuccessfulCreatePlaylist(playlist.Title), keyboards.PlaylistMenu)
}
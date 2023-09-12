package handlers

import (
	"fmt"
	"log"
	"sync"

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
	return c.Send(messages.SendPlaylistURL, keyboards.BackMenu)
}

func OnRecivePlaylistURL(c telebot.Context, s fsm.Context) error {
	if err := utils.ValidateYtURL(c.Message().Text); err != nil {
		return c.Send(messages.ErrParseYtURL)
	}
	defer finish(s)

	user := c.Get(callbacks.UserKey).(models.User)

	playlist := models.Playlist{
		URL:    c.Message().Text,
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
	if err := c.Send(playlist.String(), keyboards.PlaylistMenu); err != nil {
		log.Print(err)
	}

	videoChan := make(chan models.Video)

	go func() {
		wg := sync.WaitGroup{}
		for _, video := range videos {
			wg.Add(1)
			go func(video models.Video) error {
				defer wg.Done()
				if err := video.ParseYt(); err != nil {
					log.Print(err)
					return c.Send(messages.ErrParseYtURL)
				}
				videoChan <- video
				return nil
			}(video)
		}
		wg.Wait()
		close(videoChan)
	}()

	for v := range videoChan {
		go func(v models.Video) error {
			v.PlaylistID = playlist.ID
			if err := v.Create(); err != nil {
				log.Print(err)
				return c.Send(messages.ErrLoadVideoFromYt)
			}
			return c.Send(v.String())
		}(v)
	}

	// for _, video := range videos {
	// 	go func(c telebot.Context, video models.Video) error {
	// 		if err := video.ParseYt(); err != nil {
	// 			log.Print(err)
	// 			return c.Send(messages.ErrParseYtURL)
	// 		}
	// 		video.PlaylistID = playlist.ID
	// 		if err := video.Create(); err != nil {
	// 			log.Print(err)
	// 			return c.Send(messages.ErrLoadVideoFromYt)
	// 		}
	// 		return c.Send(video.String())
	// 	}(c, video)
	// }
	return nil
	// return c.Send(playlist.String(), keyboards.PlaylistMenu)
}

func OnEditPlaylistInlineBtn(c telebot.Context, _ fsm.Context) error {
	playlist, err := callbackWithPlaylist(c)
	if err != nil {
		return err
	}
	return c.Send(playlist.String(), keyboards.EditPlaylistInline(playlist.ID))
}

func OnShowPlaylistInlineBtn(c telebot.Context, _ fsm.Context) error {
	playlist, err := callbackWithPlaylist(c)
	if err != nil {
		return err
	}
	var message string
	for i, v := range playlist.Videos {
		message += fmt.Sprintf("%d. %s\n", i+1, v)
		if i%10 == 0 {
			if err := c.Send(message); err != nil {
				log.Print(err)
			}
			message = ""
		}
	}
	return c.Send(message)
}

func OnDeletePlaylistInlineBtn(c telebot.Context, _ fsm.Context) error {
	playlist, err := callbackWithPlaylist(c)
	if err != nil {
		return err
	}
	if err := playlist.CascadeDelete(); err != nil {
		log.Print(err)
		return c.Send(messages.ErrDeletePlaylist, keyboards.PlaylistMenu)
	}
	return c.Send(messages.SuccessfulDeletePlaylist, keyboards.PlaylistMenu)
}

func callbackWithPlaylist(c telebot.Context) (models.Playlist, error) {
	defer delMes(c)
	playlist := models.Playlist{}
	if err := playlist.Get(utils.ToUint(c.Get(callbacks.CallbackVal))); err != nil {
		log.Print(err)
		return playlist, c.Send(messages.ErrGetPlaylist, keyboards.PlaylistMenu)
	}
	return playlist, nil
}

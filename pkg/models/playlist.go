package models

import (
	// "errors"
	"fmt"
	// "log"

	"github.com/Devil666face/gotubebot/pkg/utils"
	// "github.com/Devil666face/gotubebot/pkg/config"
	"github.com/Devil666face/gotubebot/pkg/store/database"

	"gorm.io/gorm"
)

type Playlist struct {
	gorm.Model
	Title  string
	Url    string
	UserID uint
	Videos []Video
}

func (playlist Playlist) String() string {
	return fmt.Sprintf("<a href='%s'>%s</a>", playlist.Url, playlist.Title)
}

func (playlist *Playlist) Get(id uint) error {
	if err := database.DB.Preload("Videos").First(playlist, id); err != nil {
		return err.Error
	}
	return nil
}

func (playlist *Playlist) ParseYt() ([]Video, error) {
	videos := []Video{}
	title, videoUrls, err := utils.PlaylistInfo(playlist.Url)
	if err != nil {
		return videos, err
	}
	playlist.Title = title

	for _, href := range videoUrls {
		videos = append(videos, Video{
			Url:    href,
			UserID: playlist.UserID,
		})
	}
	return videos, nil
}

func (playlist *Playlist) Create() error {
	if err := database.DB.Save(playlist); err != nil {
		return err.Error
	}
	return nil
}

func GetAllPlaylistsForUser(id uint) ([]Playlist, error) {
	var playlists = []Playlist{}
	if err := database.DB.Where("user_id = ?", id).Find(&playlists); err != nil {
		return playlists, err.Error
	}
	return playlists, nil
}

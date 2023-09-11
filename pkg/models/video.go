package models

import (
	// "errors"

	"fmt"

	"github.com/Devil666face/gotubebot/pkg/utils"
	// "github.com/Devil666face/gotubebot/pkg/config"
	"github.com/Devil666face/gotubebot/pkg/store/database"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Title       string
	Url         string
	DownloadUrl string
	UserID      uint
	PlaylistID  uint
}

func (video Video) String() string {
	return fmt.Sprintf("<a href='%s'>%s</a>\n<a href='%s'>Download</a>", video.Url, video.Title, video.DownloadUrl)
}

func (video *Video) Get(id uint) error {
	if err := database.DB.First(video, id); err != nil {
		return err.Error
	}
	return nil
}

func (video *Video) Create() error {
	if err := database.DB.Save(video); err != nil {
		return err.Error
	}
	return nil
}

func (video *Video) Update() error {
	if err := database.DB.Save(video); err != nil {
		return err.Error
	}
	return nil
}

func (video *Video) Delete() error {
	if err := database.DB.Unscoped().Delete(video); err != nil {
		return err.Error
	}
	return nil
}

func (video *Video) ParseYt() error {
	title, downloadUrl, err := utils.VideoInfo(video.Url)
	if err != nil {
		return err
	}
	video.Title = title
	video.DownloadUrl = downloadUrl
	return nil
}

func GetAllVideosForUser(id uint) ([]Video, error) {
	var videos = []Video{}
	if err := database.DB.Where("user_id = ?", id).Where("playlist_id = ?", 0).Find(&videos); err != nil {
		return videos, err.Error
	}
	return videos, nil
}

package models

import (
	"fmt"

	"github.com/Devil666face/gotubebot/pkg/store/database"
	"github.com/Devil666face/gotubebot/pkg/utils"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Title       string
	URL         string
	DownloadURL string
	AudioURL    string
	UserID      uint
	PlaylistID  uint
}

func (video Video) String() string {
	return fmt.Sprintf("<a href='%s'>%s</a>\n<a href='%s'><b>🎥Download video</b></a>\n<a href='%s'><b>🎵Download audio</b></a>", video.URL, video.Title, video.DownloadURL, video.AudioURL)
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
	title, downloadURL, audioUrl, err := utils.VideoInfo(video.URL)
	if err != nil {
		return err
	}
	video.Title = title
	video.DownloadURL = downloadURL
	video.AudioURL = audioUrl
	return nil
}

func GetAllVideosForUser(id uint) ([]Video, error) {
	var videos = []Video{}
	if err := database.DB.Where("user_id = ?", id).Where("playlist_id = ?", 0).Find(&videos); err != nil {
		return videos, err.Error
	}
	return videos, nil
}

func GetAllExpireVideos() ([]Video, error) {
	var videos = []Video{}
	if err := database.DB.Where("updated_at < ?", utils.GetOneHourAgo()).Find(&videos); err != nil {
		return videos, err.Error
	}
	return videos, nil
}

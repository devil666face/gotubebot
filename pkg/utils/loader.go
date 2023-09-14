package utils

import (
	"github.com/kkdai/youtube/v2"
)

func VideoInfo(href string) (title string, downloadURL string, audioURL string, err error) {
	client := youtube.Client{}

	video, err := client.GetVideo(href)
	if err != nil {
		return
	}

	title = video.Title

	formats := video.Formats.WithAudioChannels()

	if v := formats[0]; v.URL != "" {
		downloadURL = v.URL
	}

	if v := formats[len(formats)-2]; v.URL != "" {
		audioURL = v.URL
	}
	return
}

func PlaylistInfo(href string) (string, []string, error) {
	videoUrls := []string{}
	client := youtube.Client{}

	playlist, err := client.GetPlaylist(href)
	if err != nil {
		return "", videoUrls, err
	}

	for _, v := range playlist.Videos {
		videoUrls = append(videoUrls, toYtURL(v.ID))
	}

	return playlist.Title, videoUrls, nil
}

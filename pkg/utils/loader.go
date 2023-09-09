package utils

import (
	"github.com/kkdai/youtube/v2"
)

func VideoInfo(href string) (string, string, error) {
	client := youtube.Client{}

	video, err := client.GetVideo(href)
	if err != nil {
		return "", "", err
	}

	formats := video.Formats.WithAudioChannels()

	return video.Title, formats[0].URL, nil
}

func PlaylistInfo(href string) (string, []string, error) {
	videoUrls := []string{}
	client := youtube.Client{}

	playlist, err := client.GetPlaylist(href)
	if err != nil {
		return "", videoUrls, err
	}

	for _, v := range playlist.Videos {
		videoUrls = append(videoUrls, toYtUrl(v.ID))
	}

	return playlist.Title, videoUrls, nil
}

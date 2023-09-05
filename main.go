package main

import (
	"io"
	"os"

	"github.com/kkdai/youtube/v2"
)

func main() {
	ExampleClient()
}

// ExampleDownload : Example code for how to use this package for download video.
func ExampleClient() {
	videoID := "https://www.youtube.com/watch?v=uQE7H7oI_4Q"
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		panic(err)
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}

	file, err := os.Create("video.mp4")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}
}

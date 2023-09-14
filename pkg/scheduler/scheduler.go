package scheduler

import (
	"log"
	"sync"

	"github.com/Devil666face/gotubebot/pkg/models"
)

func UpdateAllVideos() {
	videos, err := models.GetAllExpireVideos()
	if err != nil {
		log.Print(err)
		return
	}

	videoChan := make(chan models.Video)

	go func() {
		wg := sync.WaitGroup{}
		for _, video := range videos {
			wg.Add(1)
			go func(video models.Video) {
				defer wg.Done()
				if err := video.ParseYt(); err != nil {
					log.Print(err)
				}
				videoChan <- video
			}(video)
		}
		wg.Wait()
		close(videoChan)
	}()

	for v := range videoChan {
		go func(v models.Video) {
			if err := v.Update(); err != nil {
				log.Print(err)
			}
		}(v)
	}
}

package updater

import (
	"time"

	"github.com/Devil666face/gotubebot/pkg/scheduler"
)

func Start() {
	ticker := time.NewTicker(30 * time.Minute)
	scheduler.UpdateAllVideos()
	go func() {
		for {
			select {
			case <-ticker.C:
				scheduler.UpdateAllVideos()
			}
		}
	}()
}
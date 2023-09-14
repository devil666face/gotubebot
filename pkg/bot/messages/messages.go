package messages

import (
	"fmt"
)

const (
	Back   = "🔙Back"
	Update = "🔄Update"
	Delete = "🗑️Delete"
)

const (
	Videos = "🎥Videos"
)

const (
	Playlists = "🎞️Playlists"
)

func ErrSendMessage(username string) string {
	return fmt.Sprintf("❌Error to send message for user - @%s", username)
}

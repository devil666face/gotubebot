package messages

import (
	"fmt"
)

const (
	Back   = "Back"
	GoBack = "Go back"
	Update = "Update"
	Delete = "Delete"
)

const (
	Videos = "Videos"
)

const (
	Playlists = "Playlists"
)

func ErrSendMessage(username string) string {
	return fmt.Sprintf("Error to send message for user - @%s", username)
}

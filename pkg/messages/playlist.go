package messages

import "fmt"

const (
	CreatePlaylist        = "Create playlist"
	ChangePlaylist        = "Change playlist"
	PlaylistList          = "All playlists"
	SendPlaylistUrl       = "Send url to playlist"
	ErrLoadPlaylistFromYt = "Error to load playlist"
)

func SuccessfulCreatePlaylist(title string) string {
	return fmt.Sprintf("Successful create playlist: %s", title)
}

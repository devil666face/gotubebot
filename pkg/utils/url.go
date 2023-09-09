package utils

import (
	"fmt"
	"net/url"
)

const (
	YtURL = "https://youtu.be/"
)

func ValidateYtURL(href string) error {
	if _, err := url.ParseRequestURI(href); err != nil {
		return err
	}
	return nil
}

func toYtUrl(vid string) string {
	return fmt.Sprintf("%s%s", YtURL, vid)
}

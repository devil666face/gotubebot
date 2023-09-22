package utils

import (
	"fmt"
	"net/url"
	"strconv"
)

const (
	YtURL       = "https://youtu.be/"
	ExpireParam = "expire"
)

func ValidateYtURL(href string) error {
	if _, err := url.ParseRequestURI(href); err != nil {
		return err
	}
	return nil
}

func toYtURL(vid string) string {
	return fmt.Sprintf("%s%s", YtURL, vid)
}

func GetExpireParam(href string) (int64, error) {
	URL, err := url.ParseRequestURI(href)
	if err != nil {
		return 0, err
	}
	expire := URL.Query().Get(ExpireParam)
	e, err := strconv.Atoi(expire)
	if err != nil {
		return 0, err
	}
	return int64(e), nil
}

package database

import (
	"os"
	"path/filepath"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Migrate(tables ...interface{}) error {
	return DB.AutoMigrate(tables...)
}

func setPath(file string) (string, error) {
	base, err := os.Getwd()
	if err != nil {
		return "", err
	}
	abs, err := filepath.Abs(filepath.Join(base, file))
	if err != nil {
		return "", err
	}
	dir := filepath.Dir(abs)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return "", err
		}
	}
	return abs, nil
}

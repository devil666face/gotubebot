package database

import (
	"time"

	"github.com/Devil666face/gotubebot/pkg/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Sqlite3Database(cfg config.Config) error {
	path, err := setPath(cfg.Db)
	if err != nil {
		return err
	}
	if DB, err = gorm.Open(sqlite.Open(path+"?cache=shared&mode=rwc&_busy_timeout=50000"), &gorm.Config{
		NowFunc: func() time.Time { return time.Now().Local() },
		Logger:  logger.Default.LogMode(logger.Info),
	}); err != nil {
		return err
	}
	return nil
}

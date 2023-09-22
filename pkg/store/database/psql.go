package database

import (
	"fmt"
	"time"

	"github.com/Devil666face/gotubebot/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func psqlDatabase(cfg config.Config) (err error) {
	if DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow statement_timeout=0",
			cfg.PostgresHost,
			cfg.PostgresUser,
			cfg.PostgresPassword,
			cfg.PostgresDb,
			cfg.PostgresPort,
		),
	}), &gorm.Config{
		NowFunc: func() time.Time { return time.Now().Local() },
		Logger:  logger.Default.LogMode(logger.Info),
	}); err != nil {
		return err
	}
	return nil
}

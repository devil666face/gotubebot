package bot

import (
	"time"

	"github.com/Devil666face/gotubebot/pkg/config"
	"github.com/Devil666face/gotubebot/pkg/models"
	"github.com/Devil666face/gotubebot/pkg/store/database"

	telebot "gopkg.in/telebot.v3"
)

func StartBot() (*telebot.Bot, error) {
	if err := config.New(); err != nil {
		return nil, err
	}
	if err := database.Sqlite3Database(config.Cfg); err != nil {
		return nil, err
	}
	if err := database.Migrate(
		&models.User{},
	); err != nil {
		return nil, err
	}
	return New(config.Cfg)
}

func New(cfg config.Config) (*telebot.Bot, error) {
	return telebot.NewBot(telebot.Settings{
		Token:     cfg.Token,
		Poller:    &telebot.LongPoller{Timeout: 10 * time.Second},
		Verbose:   cfg.Debug,
		ParseMode: telebot.ModeHTML,
	})
}

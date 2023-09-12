package bot

import (
	"time"

	"github.com/Devil666face/gotubebot/pkg/bot/routes"
	"github.com/Devil666face/gotubebot/pkg/config"
	"github.com/Devil666face/gotubebot/pkg/models"
	"github.com/Devil666face/gotubebot/pkg/store/database"
	"github.com/Devil666face/gotubebot/pkg/store/mem"
	"github.com/vitaliy-ukiru/fsm-telebot"

	telebot "gopkg.in/telebot.v3"
)

func Get() (*telebot.Bot, error) {
	if err := config.New(); err != nil {
		return nil, err
	}
	if err := database.Sqlite3Database(config.Cfg); err != nil {
		return nil, err
	}
	if err := database.Migrate(
		&models.Video{},
		&models.Playlist{},
		&models.User{},
	); err != nil {
		return nil, err
	}
	bot, err := newBot()
	if err != nil {
		return nil, err
	}
	manager := newFsmManager(bot)
	routes.New(manager)
	return bot, nil
}

func newBot() (*telebot.Bot, error) {
	return telebot.NewBot(telebot.Settings{
		Token:     config.Cfg.Token,
		Poller:    &telebot.LongPoller{Timeout: 10 * time.Second},
		Verbose:   config.Cfg.Debug,
		ParseMode: telebot.ModeHTML,
	})
}

func newFsmManager(bot *telebot.Bot) *routes.Manager {
	return &routes.Manager{
		fsm.NewManager(bot, nil, mem.New(), nil),
	}
}

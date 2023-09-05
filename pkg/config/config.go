package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

var Cfg Config

type Config struct {
	Token       string `env:"TOKEN" env-required:"true"`
	Db          string `env:"DB" env-default:"db.sqlite3"`
	Debug       bool   `env:"DEBUG" env-default:"false"`
	SuperuserID uint   `env:"SUPERUSERID" env-default:"446545799"`
	Log         bool   `env:"LOG" env-default:"true"`
}

func New() error {
	return cleanenv.ReadEnv(&Cfg)
}

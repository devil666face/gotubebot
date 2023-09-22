package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

var Cfg Config

type Config struct {
	Token            string `env:"TOKEN" env-required:"true"`
	Db               string `env:"DB" env-default:"db.sqlite3"`
	Debug            bool   `env:"DEBUG" env-default:"false"`
	SuperuserID      uint   `env:"SUPERUSERID" env-default:"446545799"`
	Log              bool   `env:"LOG" env-default:"true"`
	PostgresUse      bool   `env:"POSTGRES" env-default:"false"`
	PostgresDb       string `env:"POSTGRES_DB" env-default:"db"`
	PostgresUser     string `env:"POSTGRES_USER" env-default:"superuser"`
	PostgresPassword string `env:"POSTGRES_PASSWORD" env-default:"Qwerty123!"`
	PostgresHost     string `env:"POSTGRES_HOST" env-default:"localhost"`
	PostgresPort     string `env:"POSTGRES_PORT" env-default:"5432"`
}

func New() error {
	return cleanenv.ReadEnv(&Cfg)
}

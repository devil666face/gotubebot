package mem

import (
	"github.com/vitaliy-ukiru/fsm-telebot/storages/memory" //nolint:misspell
)

func New() *memory.Storage {
	return memory.NewStorage()
}

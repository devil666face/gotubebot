package mem

import (
	"github.com/vitaliy-ukiru/fsm-telebot/storages/memory"
)

func New() *memory.Storage {
	return memory.NewStorage()
}

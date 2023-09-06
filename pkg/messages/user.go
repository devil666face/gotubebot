package messages

import (
	"fmt"

	"gopkg.in/telebot.v3"
)

const (
	ConfirmUser     = "Add this user"
	IgnoreUser      = "Ignore this user"
	ErrGetAdminList = "Failed to get admins list"
)

func AskAdminsForAddUser(c telebot.Context) string {
	return fmt.Sprintf("Добавить пользователя @%s - %d?", c.Chat().Username, c.Chat().ID)
}

func SuccessfulCreateUser(c telebot.Context) string {
	return fmt.Sprintf("Администратору отправлен запрос на создание вашего пользователя - @%s", c.Chat().Username)
}

func SuccessfulUpdateUser(username string) string {
	return fmt.Sprintf("Пользователь @%s - успешно обновлен", username)
}

func PermissionsForUserAdded(username string) string {
	return fmt.Sprintf("Администратор добавил вас - @%s", username)
}

func ErrCreateUser(c telebot.Context) string {
	return fmt.Sprintf("Произошла ошибка при создании пользователя - @%s", c.Chat().Username)
}

func ErrUserNotFound(id int64) string {
	return fmt.Sprintf("Пользователь с id - %d не найден", id)
}

func ErrUserUpdate(username string) string {
	return fmt.Sprintf("Ошибка обновления пользователя @%s", username)
}

func ErrUserAlreadyCreate(c telebot.Context) string {
	return fmt.Sprintf("Пользователь @%s - уже создан", c.Chat().Username)
}

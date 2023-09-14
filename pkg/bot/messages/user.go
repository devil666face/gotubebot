package messages

import (
	"fmt"

	"gopkg.in/telebot.v3"
)

const (
	ConfirmUser     = "❔Add this user"
	IgnoreUser      = "⛔️Ignore this user"
	ErrGetAdminList = "❌Failed to get admins list"
	ErrGetUser      = "❌Failed to get ower user"
)

func AskAdminsForAddUser(c telebot.Context) string {
	return fmt.Sprintf("❔Add user @%s - %d?", c.Chat().Username, c.Chat().ID)
}

func SuccessfulCreateUser(c telebot.Context) string {
	return fmt.Sprintf("❔An administrator has been sent a request to create your user - @%s", c.Chat().Username)
}

func SuccessfulUpdateUser(username string) string {
	return fmt.Sprintf("✅User @%s - successful update", username)
}

func PermissionsForUserAdded(username string) string {
	return fmt.Sprintf("✅Administrator added you - @%s", username)
}

func ErrCreateUser(c telebot.Context) string {
	return fmt.Sprintf("❌Error when create user - @%s", c.Chat().Username)
}

func ErrUserNotFound(id int64) string {
	return fmt.Sprintf("❌User id - %d not found", id)
}

func ErrUserUpdate(username string) string {
	return fmt.Sprintf("❌Error update user @%s", username)
}

func ErrUserAlreadyCreate(c telebot.Context) string {
	return fmt.Sprintf("❌User @%s - already create", c.Chat().Username)
}

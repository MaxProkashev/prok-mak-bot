package logic

import (
	"database/sql"
	dbfunc "prok-mak-bot/pkg/db-func"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// HookConfig - information about telegram request by user
type HookConfig struct {
	UpdateID int   // id req
	ChatID   int64 // id chat
	UserID   int   // id user

	HasText     bool // is there any text
	HasPhoto    bool // is there any photo
	HasCallback bool // is there any callback

	InTable bool // was the user in the table already
}

// ParseUpdate - get main info about req
func ParseUpdate(db *sql.DB, update tgbotapi.Update) (hook HookConfig) {
	if update.CallbackQuery != nil {
		hook = HookConfig{
			HasCallback: true,
			ChatID:      update.CallbackQuery.Message.Chat.ID,
			UserID:      update.CallbackQuery.From.ID,
		}
	} else if update.Message != nil {
		hook = HookConfig{
			ChatID: update.Message.Chat.ID,
			UserID: update.Message.From.ID,
		}
		if update.Message.Photo != nil {
			hook.HasPhoto = true
		}
		if update.Message.Text != "" {
			hook.HasText = true
		}
	}
	hook.UpdateID = update.UpdateID
	hook.InTable = dbfunc.CheckUserID(db, hook.UserID)
	return hook
}

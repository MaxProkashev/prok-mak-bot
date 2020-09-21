package logic

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// HookConfig - information about telegram request by user
type HookConfig struct {
	updateID int   // id req
	chatID   int64 // id chat
	userID   int   // id user

	hasText     bool // is there any text
	hasPhoto    bool // is there any photo
	hasCallback bool // is there any callback

	inTable bool // is there a user in the table
}

// ParseUpdate - get main info about req
func ParseUpdate(update tgbotapi.Update) (hook HookConfig) {
	if update.CallbackQuery != nil {
		hook = HookConfig{
			updateID:    update.UpdateID,
			hasCallback: true,
			chatID:      update.CallbackQuery.Message.Chat.ID,
			userID:      update.CallbackQuery.From.ID,
		}
		return hook
	} else if update.Message != nil {
		hook = HookConfig{
			updateID:    update.UpdateID,
			chatID: update.Message.Chat.ID,
			userID: update.Message.From.ID,
		}
		if update.Message.Photo != nil {
			hook.hasPhoto = true
			return hook
		}
		hook.hasText = true
		return hook
	}

	hook.updateID = update.UpdateID
	hook.chatID = update.Message.Chat.ID
	hook.userID = update.Message.From.ID
	return hook
}

package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// CreateBot - create bot with token by env heroku and set webhook
func CreateBot() (bot *tgbotapi.BotAPI, err error) {
	// birth bot
	bot, err = tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Printf("[X] Could not create bot. Reason: %s", err.Error())
		return nil, err
	}

	// set webhook
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(os.Getenv("COMMUNICATION_URL") + os.Getenv("TOKEN")))
	if err != nil {
		log.Printf("[X] Could not set webhook to bot settings. Reason: %s", err.Error())
		return nil, err
	}
	return bot, nil
}

// WebhookHandler - get info about webhook from tg-bot
func WebhookHandler(c *gin.Context, bot *tgbotapi.BotAPI) {
	defer c.Request.Body.Close()

	// read request
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalf("[X] Could not read request. Reason: %s", err.Error())
		return
	}

	// unmarshal update
	var update tgbotapi.Update
	err = json.Unmarshal(bytes, &update)
	if err != nil {
		log.Fatalf("[X] Could not unmarshal updates. Reason: %s", err.Error())
		return
	}

}

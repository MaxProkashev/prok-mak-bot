package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/heroku/x/hmetrics/onload"
)

var bot *tgbotapi.BotAPI

func main() {
	// get port heroku env
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("[X] $PORT must be set")
	} else {
		log.Printf("[OK] Get PORT = %s", port)
	}

	// birth of a bot
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Fatalf("[X] Could not create bot. Reason: %s", err.Error())
		return
	}

	// set webhook
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(os.Getenv("COMMUNICATION_URL") + bot.Token))
	if err != nil {
		log.Fatalf("[X] Could not set webhook to bot settings. Reason: %s", err.Error())
		return
	}

	// run gin router
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	router.POST("/"+bot.Token, webhookHandler)

	// run router
	err = router.Run(":" + port)
	if err != nil {
		log.Fatalf("[X] Could not run router. Reason: %s", err.Error())
	}
}

func webhookHandler(c *gin.Context) {
	defer c.Request.Body.Close()

	// Чтение запроса
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalf("[X] Could not read request. Reason: %s", err.Error())
		return
	}

	// Update
	var update tgbotapi.Update
	err = json.Unmarshal(bytes, &update)
	if err != nil {
		log.Fatalf("[X] Could not unmarshal updates. Reason: %s", err.Error())
		return
	}

	bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "ПРИВЕТ"))
}

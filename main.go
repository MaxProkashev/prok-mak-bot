package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	// get port heroku env
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("[X] $PORT must be set")
	} else {
		log.Printf("[OK] Get PORT = %s", port)
	}
	// run gin router
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// birth of a bot
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Fatalf("[X] Could not create bot. Reason: %s", err.Error())
		return
	}

	// communication with the server
	url := os.Getenv("COMMUNICATION_URL") + bot.Token
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(url))
	if err != nil {
		log.Fatalf("[X] Could not set webhook to bot settings. Reason: %s", err.Error())
	}

	router.POST("/" + bot.Token)

	// run router
	err = router.Run(":" + port)
	if err != nil {
		log.Fatalf("[X] Could not run router. Reason: %s", err.Error())
	}
}

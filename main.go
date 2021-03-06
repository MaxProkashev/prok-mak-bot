package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	logic "prok-mak-bot/pkg/bot-logic"
	dbfunc "prok-mak-bot/pkg/db-func"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	// get port by heroku env
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("[X] $PORT must be set")
	} else {
		log.Printf("[OK] Get PORT = %s", port)
	}

	// customize gin router
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	// create db postresql
	db, err := dbfunc.OpenDB()
	if err != nil {
		log.Fatalf("[X] Could not connect to DB. Reason: %s", err.Error())
	} else {
		log.Printf("[OK] Connect to DB")
	}
	dbfunc.CreateTable(db)

	// create bot
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Fatalf("[X] Could not create bot. Reason: %s", err.Error())
	}

	// set webhook
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(os.Getenv("COMMUNICATION_URL") + os.Getenv("TOKEN")))
	if err != nil {
		log.Fatalf("[X] Could not set webhook to bot settings. Reason: %s", err.Error())
	}

	// processing request
	router.POST("/"+os.Getenv("TOKEN"), func(c *gin.Context) {
		defer c.Request.Body.Close()

		// read request body
		bytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Fatalf("[X] Could not read request. Reason: %s", err.Error())
		}

		// unmarshal update
		var update tgbotapi.Update
		err = json.Unmarshal(bytes, &update)
		if err != nil {
			log.Fatalf("[X] Could not unmarshal updates. Reason: %s", err.Error())
		}

		hook := logic.ParseUpdate(db, update)
		if hook.InTable {
			bot.Send(tgbotapi.NewMessage(hook.ChatID, update.Message.From.UserName))
		} else {
			bot.Send(tgbotapi.NewMessage(hook.ChatID, "Вы внесены в таблицу"))
		}
	})

	// run gin router
	err = router.Run(":" + port)
	if err != nil {
		log.Fatalf("[X] Could not run router. Reason: %s", err.Error())
	}
}

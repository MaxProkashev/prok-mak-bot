package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"

	"prok-mak-bot/pkg/handler"
)

func main() {
	// get port by heroku env
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("[X] $PORT must be set")
	} else {
		log.Printf("[OK] Get PORT = %s", port)
	}

	// run gin router
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	// create bot
	bot, err := handler.CreateBot()
	if err != nil {
		log.Fatal("[X] Could not create bot")
	}

	router.POST("/"+os.Getenv("TOKEN"), func(c *gin.Context) {
		log.Println(c)
		log.Println(bot)
		handler.WebhookHandler(c, bot)
	})

	// run router
	err = router.Run(":" + port)
	if err != nil {
		log.Fatalf("[X] Could not run router. Reason: %s", err.Error())
	}
}

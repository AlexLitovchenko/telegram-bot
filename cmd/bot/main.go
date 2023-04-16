package main

import (
	"os"

	"github.com/AlexLitovchenko/telegram-bot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error initializing env variable: %s", err.Error())
	}
	var (
		myToken = os.Getenv("myToken")
		//myID    = os.Getenv("myID")
		//groupID = os.Getenv("groupID")
	)
	log.Info(myToken)
	bot, err := tgbotapi.NewBotAPI(myToken)
	if err != nil {
		log.Panicf("New bot error: %v", err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot)

	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}

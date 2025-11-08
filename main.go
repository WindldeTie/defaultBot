package main

import (
	"DeliveryBot_mini/handler"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	godotenv.Load()
	bot := initBot()
	bot.Request(tgbotapi.DeleteWebhookConfig{DropPendingUpdates: true})

	handler.NewHandler(bot).Start(false)
}

func initBot() *tgbotapi.BotAPI {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN environment variable is required")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panicf("Bot creation failed: %v", err)
	}

	log.Printf("✅ Authorized as @%s", bot.Self.UserName)
	return bot
}

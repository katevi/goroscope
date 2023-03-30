package main

import (
	"context"
	"log"
	"os"

	generator "goroscope/internal"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

const (
	tokenKey          = "TELEGRAM_BOT_TOKEN"
	startCommand      = "/start"
	goroscopeCommand  = "/goroscope"
	envFileExt        = ".env"
	timeoutForUpdates = 60
)

func main() {
	bot := configureBot()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = timeoutForUpdates
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		log.Printf("A message %s was received from %v", update.Message.Text, update.Message.From)
		msg := handleUpdate(context.Background(), update)
		bot.Send(msg)
	}
}

func handleUpdate(ctx context.Context, update tgbotapi.Update) tgbotapi.MessageConfig {
	switch update.Message.Text {
	case startCommand:
		greetingMsg := "Good morning to your majesty."
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, greetingMsg)
		return msg
	case goroscopeCommand:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, generator.GenerateHoroscope())
		return msg
	default:
		repeatMsg := "Could you please repeat your wisdom, sir."
		return tgbotapi.NewMessage(update.Message.Chat.ID, repeatMsg)
	}
}

func configureBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(getToken(tokenKey))
	if err != nil {
		log.Fatalf("Error occurred during bot initialization %s ", err)
	}
	addCommands(bot)
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	return bot
}

func addCommands(bot *tgbotapi.BotAPI) {
	myCommandsConfig := tgbotapi.SetMyCommandsConfig{Commands: []tgbotapi.BotCommand{
		{Command: startCommand, Description: "get hello and introduction from bot"},
		{Command: goroscopeCommand, Description: "get goroscope"},
	}, LanguageCode: "en"}

	resp, err := bot.Request(myCommandsConfig)
	if err != nil {
		log.Printf("Failed to receive response from bot %v", resp)
	}
}

func getToken(key string) string {
	err := godotenv.Load(envFileExt)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

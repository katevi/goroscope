package main

import (
	"context"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"goroscope.bot/horoscope"
)

const (
	TELEGRAM_BOT_TOKEN = "TELEGRAM_BOT_TOKEN"
)

func main() {
	bot := configureBot()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		log.Printf("A message %s was received from %v", update.Message.Text, update.Message.From)
		msg := makeServiceHandler(context.Background(), update)
		bot.Send(msg)
	}
}

func makeServiceHandler(ctx context.Context, update tgbotapi.Update) tgbotapi.MessageConfig {
	switch update.Message.Text {
	case "/start":
		greetingMsg := "Good morning to your majesty."
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, greetingMsg)
		msg.ReplyToMessageID = update.Message.MessageID
		return msg
	case "/goroscope":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, horoscope.GenerateHoroscope())
		msg.ReplyToMessageID = update.Message.MessageID
		return msg
	default:
		repeatMsg := "Could you please repeat your wisdom, sir."
		return tgbotapi.NewMessage(update.Message.Chat.ID, repeatMsg)
	}
}

func configureBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(getToken(TELEGRAM_BOT_TOKEN))
	if err != nil {
		log.Fatalf("Error occurred during bot initialization %s ", err)
	}
	addMyCommands(bot)
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	return bot
}

func getToken(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func addMyCommands(bot *tgbotapi.BotAPI) {
	myCommandsConfig := tgbotapi.SetMyCommandsConfig{Commands: []tgbotapi.BotCommand{
		{Command: "/start", Description: "get introduction"},
		{Command: "/goroscope", Description: "get goroscope"},
	}, LanguageCode: "en"}

	resp, err := bot.Request(myCommandsConfig)
	if err != nil {
		log.Printf("Failed to receive response from bot %v", resp)
	}
}

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	generator "goroscope/internal"
	"goroscope/internal/store"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

const (
	tokenKey           = "TELEGRAM_BOT_TOKEN"
	startCommand       = "/start"
	goroscopeCommand   = "/goroscope"
	subscribeCommand   = "/subscribe"
	unsubscribeCommand = "/unsubscribe"
	envFileExt         = ".env"
	timeoutForUpdates  = 60
	period             = 24
)

var help = strings.Join([]string{
	"I can give you goroscope in \"Heroes of Might and Magic\" game style. ",
	fmt.Sprintf("Please, enter %s to get your goroscope.", goroscopeCommand),
	fmt.Sprintf("Please, enter %s to get new goroscope every 24h.", subscribeCommand),
	fmt.Sprintf("Please, enter %s to stop getting new goroscopes every 24h.", unsubscribeCommand),
}, "\n")

func main() {
	bot := configureBot()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = timeoutForUpdates
	updates := bot.GetUpdatesChan(u)
	subscribers := store.NewSubscribers()

	go sendGoroscopePeriodically(bot, &subscribers)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Text != "" {
			log.Printf("A message %s was received from %v", update.Message.Text, update.Message.From)
			msg := handleUpdate(context.Background(), update, subscribers)
			bot.Send(msg)
		}
	}
}

func sendGoroscopePeriodically(bot *tgbotapi.BotAPI, subscribers *store.Subscribers) {
	ticker := time.NewTicker(period * time.Hour)
	for range ticker.C {
		log.Printf("Time ticked. Subscribers = %v\n", subscribers.All())
		for _, subscriber := range subscribers.All() {
			msg := tgbotapi.NewMessage(subscriber, generator.GenerateHoroscope())
			bot.Send(msg)
		}
	}
}

func handleUpdate(ctx context.Context, update tgbotapi.Update, subscribers store.Subscribers) tgbotapi.MessageConfig {
	switch update.Message.Text {
	case startCommand:
		greetingMsg := fmt.Sprint("Good day to your majesty.\n", help)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, greetingMsg)
		return msg
	case goroscopeCommand:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, generator.GenerateHoroscope())
		return msg
	case subscribeCommand:
		subscribers.Add(update.Message.Chat.ID)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Subscribed successfully! Your goroscope will be delivered within next 24h =)")
		return msg
	case unsubscribeCommand:
		subscribers.Rm(update.Message.Chat.ID)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Unsubscribed successfully.")
		return msg
	default:
		repeatMsg := fmt.Sprint("Could you please repeat your wisdom, sir!\n", help)
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
		{Command: subscribeCommand, Description: "get goroscope every day"},
		{Command: unsubscribeCommand, Description: "stop getting goroscope every day"},
	}, LanguageCode: "en"}

	resp, err := bot.Request(myCommandsConfig)
	if err != nil {
		log.Printf("Failed to receive response from bot %v", resp)
	}
}

func getToken(key string) string {
	err := godotenv.Load(envFileExt)
	if err != nil {
		log.Print("Error loading .env file")
	}
	return os.Getenv(key)
}

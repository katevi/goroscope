package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

const (
	TELEGRAM_BOT_TOKEN = "TELEGRAM_BOT_TOKEN"
)

func getToken(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	bot, err := tgbotapi.NewBotAPI(getToken(TELEGRAM_BOT_TOKEN))
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("A message %s was received from %v", update.Message.Text, update.Message.From)
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, generateHoroscope())
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

// Goroscope generation

var (
	creatures = []string{"Red Dragon", "Black dragon", "Peasant",
		"Goblin", "Dwarf", "Orc", "Gargoyle", "Ghost", "Genie", "Medusa",
		"Halfling", "Boar", "Iron Golem", "Roc", "Orc", "Wolf", "Ogre", "Troll", "Cyclop",
		"Unicorn", "Phoenix", "Sprite", "Dwarf"}
	plagueWeek        = "Plague"
	weekOf            = append(creatures, plagueWeek)
	necromances       = []string{"Zombie", "Skeleton", "Mummy", "Vampire", "Lich"}
	dwellingsSentence = "All dwellings increase population."
)

func generateHoroscope() string {
	fst, seed := generateFirstSentence()
	snd := generateSecondSentence(seed)
	return fmt.Sprint(fst, snd)
}

func generateFirstSentence() (sentence, seed string) {
	weekOf := pickRandomElement(weekOf)
	return fmt.Sprintf("Astrologers proclaim the week of %s. ", weekOf), weekOf
}

func generateSecondSentence(seed string) string {
	if seed == plagueWeek {
		return generateNecromancesIncreased()
	}
	return generateCreaturesIncreased(seed)
}

func generateNecromancesIncreased() string {
	return fmt.Sprintf("%s growth +%d. ", pickRandomElement(necromances), rand.Intn(10))
}

func generateCreaturesIncreased(seed string) string {
	creaturesIncreased := fmt.Sprintf("%s growth +%d. ", seed, rand.Intn(10))
	if randBool() {
		return fmt.Sprint(creaturesIncreased, dwellingsSentence)
	}
	return creaturesIncreased
}

// Helper functions for generation

func randBool() bool {
	return rand.Intn(2) == 1
}

func pickRandomElement(array []string) string {
	randomIndex := rand.Intn(len(array))
	return array[randomIndex]
}

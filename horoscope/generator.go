package horoscope

import (
	"fmt"
	"math/rand"
)

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

func GenerateHoroscope() string {
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

package horoscope

import (
	"fmt"
	"math/rand"
)

// Goroscope generation

var (
	creatures = []string{"Red Dragon", "Black dragon", "Peasant", "Goblin", "Dwarf",
		"Orc", "Gargoyle", "Ghost", "Genie", "Medusa",
		"Halfling", "Boar", "Iron Golem", "Roc", "Orc",
		"Wolf", "Ogre", "Troll", "Cyclop", "Unicorn",
		"Phoenix", "Sprite", "Dwarf", "Druid", "Eagle",
		"Druid", "Elf", "Hydra", "Minotaur", "Griffin"}
	goodWeek          = []string{"Health", "Luck", "Leadership", "Squirrel", "Harvest", "Good weather", "Peace", "Celebration"}
	plagueWeek        = []string{"Plague", "Epidemic", "Bad weather", "Drought", "Starvation"}
	necromances       = []string{"Zombie", "Skeleton", "Mummy", "Vampire", "Lich", "Bone Dragon"}
	dwellingsSentence = "All dwellings increase population."
	multiples         = []string{"doubles", "triples", "quadruples"}
	decreased         = []string{"halved", "decreased", "reduced"}
)

func GenerateHoroscope() string {
	return generateHoroscope(rollDice())
}

func generateHoroscope(badWeek bool) string {
	if badWeek {
		fst := generateBadWeek()
		snd := generateBadWeekPopulation()
		return fmt.Sprint(fst, snd)
	}
	fst, creature := generateGoodWeek()
	snd := generateGoodWeekPopulation(creature)
	return fmt.Sprint(fst, snd)
}

func generateBadWeek() string {
	badWeek := pickRandomElement(plagueWeek)
	return fmt.Sprintf("Astrologers proclaim the week of %s. ", badWeek)
}

func generateGoodWeek() (string, string) {
	if tossCoin() {
		goodWeek := pickRandomElement(goodWeek)
		return fmt.Sprintf("Astrologers proclaim the week of %s. ", goodWeek), pickRandomElement(creatures)
	}
	creatureWeek := pickRandomElement(creatures)
	return fmt.Sprintf("Astrologers proclaim the week of %s. ", creatureWeek), creatureWeek
}

func generateBadWeekPopulation() string {
	if tossCoin() {
		return generateNecromancesIncreased()
	}
	return generateCreaturesDecreased()
}

func generateNecromancesIncreased() string {
	if tossCoin() {
		return fmt.Sprintf("%s population %s. ", pickRandomElement(necromances), pickRandomElement(multiples))
	}
	return fmt.Sprintf("%s growth +%d. ", pickRandomElement(necromances), rand.Intn(10))
}

func generateGoodWeekPopulation(seed string) string {
	creaturesIncreased := fmt.Sprintf("%s growth +%d. ", seed, rand.Intn(10))
	if tossCoin() {
		return fmt.Sprint(creaturesIncreased, dwellingsSentence)
	}
	if tossCoin() {
		return creaturesIncreased
	}
	creaturesMultipled := fmt.Sprintf("%s population %s.", seed, pickRandomElement(multiples))
	if tossCoin() {
		return fmt.Sprint(creaturesMultipled, dwellingsSentence)
	}
	return creaturesMultipled
}

func generateCreaturesDecreased() string {
	return fmt.Sprintf("%s population %s. ", pickRandomElement(creatures), pickRandomElement(decreased))
}

// Helper functions for generation

func rollDice() bool {
	return rand.Intn(6) == 0
}

func tossCoin() bool {
	return rand.Intn(2) == 1
}

func pickRandomElement(array []string) string {
	randomIndex := rand.Intn(len(array))
	return array[randomIndex]
}

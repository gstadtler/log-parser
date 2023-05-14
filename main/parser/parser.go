package parser

import (
	"regexp"
	"strconv"
	"strings"
)

func ProcessHeal(line string) int {
	healRegex := regexp.MustCompile(`You healed yourself for (\d+) hitpoints\.`)

	resultHeal := 0

	if healRegex.MatchString(line) {
		heal, _ := strconv.Atoi(healRegex.FindStringSubmatch(line)[1])
		resultHeal = heal
	}

	return resultHeal
}

func ProcessDamage(line string) int {
	damageRegex := regexp.MustCompile(`You lose (\d+) hitpoints.`)

	resultDamage := 0

	if damageRegex.MatchString(line) {
		damage, _ := strconv.Atoi(damageRegex.FindStringSubmatch(line)[1])
		resultDamage = damage
	}

	return resultDamage
}

func ProcessCreatureDamage(line string) (string, int) {
	creatureDamageRegex := regexp.MustCompile(`You lose (\d+) hitpoints due to an attack by a (\w+)[.\s]`)

	resultCreature := ""
	resultDamage := 0

	if creatureDamageRegex.MatchString(line) {
		match := creatureDamageRegex.FindStringSubmatch(line)
		damage, _ := strconv.Atoi(match[1])
		creature := match[2]
		resultCreature = creature
		resultDamage = damage
	}

	return resultCreature, resultDamage
}

func ProcessExp(line string) int {
	expRegex := regexp.MustCompile(`You gained (\d+) experience points\.`)

	resultExp := 0

	if expRegex.MatchString(line) {
		exp, _ := strconv.Atoi(expRegex.FindStringSubmatch(line)[1])
		resultExp = exp
	}

	return resultExp
}

func ProcessLoot(line string) (string, map[string]int) {
	lootRegex := regexp.MustCompile(`Loot of a (.*): (.*)\.`)
	itemQuantityRegex := regexp.MustCompile(`(\d+) (\w+)`)

	resultCreature := ""
	resultItems := make(map[string]int)

	if lootRegex.MatchString(line) {
		match := lootRegex.FindStringSubmatch(line)

		if len(match) == 0 {
			return resultCreature, resultItems
		}

		resultCreature = match[1]
		itemsText := match[2]

		for _, item := range strings.Split(itemsText, ", ") {
			if item != "nothing" {
				var itemName string
				var itemQuantity int
				
				if submatch := itemQuantityRegex.FindStringSubmatch(item); len(submatch) > 0 {
					itemName = submatch[2]
					itemQuantity, _ = strconv.Atoi(submatch[1])
				} else {
					itemName = item[2:]
					itemQuantity = 1
				}

				resultItems[itemName] += itemQuantity
			}
		}
	}

	return resultCreature, resultItems

}

func ProcessBlackKnightDamage(line string) int {
	bkDamageRegex := regexp.MustCompile(`A Black Knight loses (\d+) hitpoints due to your attack\.`)

	resultBkDamage := 0

	if bkDamageRegex.MatchString(line) {
			damage, _ := strconv.Atoi(bkDamageRegex.FindStringSubmatch(line)[1])
			resultBkDamage = damage
		}

	return resultBkDamage
}

func ProcessUnknownDamage(line string) int {
	damageRegex := regexp.MustCompile(`You lose (\d+) hitpoints.`)
	creatureDamageRegex := regexp.MustCompile(`You lose (\d+) hitpoints due to an attack by a (\w+)[\.\s]`)

	resultUnkownDamage := 0

	if damageRegex.MatchString(line) && !creatureDamageRegex.MatchString(line) {
		damage, _ := strconv.Atoi(damageRegex.FindStringSubmatch(line)[1])
		resultUnkownDamage = damage
	}

	return resultUnkownDamage
}

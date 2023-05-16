package parser

import (
	"regexp"
	"strconv"
	"strings"
)

type Event struct {
	Description string
	Type string
	Value int
}

type DamageTaken struct {
	Total int
	PerCreatureKind map[string]int
	ByUnknownOriginTotal int
}

type Stats struct {
	DamageHealed int
	DamageTaken DamageTaken
	ExperienceGained int
	Loot map[string]map[string]int
	BlackKnightTotalHealth int
}

func ProcessHeal(line string, stats Stats, eventLog []Event) (Stats, []Event) {

	healRegex := regexp.MustCompile(`You healed yourself for (\d+) hitpoints\.`)

	if !healRegex.MatchString(line) {
		return stats, eventLog
	}

	heal, _ := strconv.Atoi(healRegex.FindStringSubmatch(line)[1])
	var event = Event {
		Description: line,
		Type: "heal",
		Value: heal,
	}

	stats.DamageHealed += heal
	eventLog = append(eventLog, event)

	return stats, eventLog
}

func ProcessDamage(line string, stats Stats, eventLog []Event) (Stats, []Event) {

	damageRegex := regexp.MustCompile(`You lose (\d+) hitpoints.`)

	if !damageRegex.MatchString(line) {
		return stats, eventLog
	}

	damage, _ := strconv.Atoi(damageRegex.FindStringSubmatch(line)[1])
	var event = Event {
		Description: line,
		Type: "damage",
		Value: damage,
	}

	stats.DamageTaken.Total += damage
	eventLog = append(eventLog, event)
	
	return stats, eventLog
}

func ProcessCreatureDamage(line string, stats Stats) Stats {

	creatureDamageRegex := regexp.MustCompile(`You lose (\d+) hitpoints due to an attack by a (\w+\s*\w*)[.\s]`)

	if !creatureDamageRegex.MatchString(line) {
		return stats
	}

	match := creatureDamageRegex.FindStringSubmatch(line)
	damage, _ := strconv.Atoi(match[1])
	creature := match[2]

	stats.DamageTaken.PerCreatureKind[creature] += damage

	return stats
}

func ProcessExp(line string, stats Stats, eventLog []Event) (Stats, []Event) {

	expRegex := regexp.MustCompile(`You gained (\d+) experience points\.`)

	if !expRegex.MatchString(line) {
		return stats, eventLog
	}

	exp, _ := strconv.Atoi(expRegex.FindStringSubmatch(line)[1])
	var event = Event {
		Description: line,
		Type: "experience",
		Value: exp,
	}

	stats.ExperienceGained += exp
	eventLog = append(eventLog, event)

	return stats, eventLog
}

func ProcessLoot(line string, stats Stats) Stats {

	lootRegex := regexp.MustCompile(`Loot of a (.*): (.*)\.`)
	itemQuantityRegex := regexp.MustCompile(`(\d+) (\w+)`)

	if !lootRegex.MatchString(line) {
		return stats
	}

	match := lootRegex.FindStringSubmatch(line)

	if len(match) == 0 {
		return stats
	}

	lootItems := make(map[string]int)
	creature := match[1]
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

			lootItems[itemName] += itemQuantity
		}
	}

	stats.Loot[creature] = lootItems

	return stats
}

func ProcessBlackKnightDamage(line string, stats Stats) Stats {

	bkDamageRegex := regexp.MustCompile(`A Black Knight loses (\d+) hitpoints due to your attack\.`)

	if !bkDamageRegex.MatchString(line) {
		return stats
	}

	damage, _ := strconv.Atoi(bkDamageRegex.FindStringSubmatch(line)[1])
	stats.BlackKnightTotalHealth += damage

	return stats
}

func ProcessUnknownDamage(line string, stats Stats) Stats {

	damageRegex := regexp.MustCompile(`You lose (\d+) hitpoints.`)
	creatureDamageRegex := regexp.MustCompile(`You lose (\d+) hitpoints due to an attack by a (\w+)[\.\s]`)

	if damageRegex.MatchString(line) && !creatureDamageRegex.MatchString(line) {
		damage, _ := strconv.Atoi(damageRegex.FindStringSubmatch(line)[1])
		stats.DamageTaken.ByUnknownOriginTotal += damage
	
		return stats
	}

	return stats
}

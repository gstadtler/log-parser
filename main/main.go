package main

import (
	"bufio"
	"fmt"
	"os"
	"encoding/json"
	"main/parser"
)

type CreatureLoot struct {
	Creature string
	Items map[string]int
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

func main() {
	file, err := os.Open("./server-log.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var stats = Stats {
		DamageTaken: DamageTaken {
			PerCreatureKind: make(map[string]int),
		},
		Loot: make(map[string]map[string]int),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		stats.DamageHealed += parser.ProcessHeal(line)

		stats.DamageTaken.Total += parser.ProcessDamage(line)

		var creature, creatureDamage = parser.ProcessCreatureDamage(line)
		stats.DamageTaken.PerCreatureKind[creature] += creatureDamage

		stats.ExperienceGained += parser.ProcessExp(line)

		var creatureName, lootItems = parser.ProcessLoot(line)
		stats.Loot[creatureName] = lootItems

		stats.BlackKnightTotalHealth += parser.ProcessBlackKnightDamage(line)

		stats.DamageTaken.ByUnknownOriginTotal += parser.ProcessUnknownDamage(line)
	}

	jsonData, err := json.Marshal(stats)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))
}
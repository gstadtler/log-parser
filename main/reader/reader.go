package reader

import (
	"bufio"
	"os"
	"main/parser"
)

var statsData = parser.Stats {
	DamageTaken: parser.DamageTaken {
		PerCreatureKind: make(map[string]int),
	},
	Loot: make(map[string]map[string]int),
}

var eventLog = []parser.Event{}

func LoadFile() {
	file, err := os.Open("./server-log.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	processFile(file)
}

func processFile(file *os.File) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		
		statsData, eventLog = parser.ProcessHeal(line, statsData, eventLog)

		statsData, eventLog = parser.ProcessDamage(line, statsData, eventLog)

		statsData = parser.ProcessCreatureDamage(line, statsData)
	
		statsData, eventLog = parser.ProcessExp(line, statsData, eventLog)
	
		statsData = parser.ProcessLoot(line, statsData)
	
		statsData = parser.ProcessBlackKnightDamage(line, statsData)
	
		statsData = parser.ProcessUnknownDamage(line, statsData)		
	}
}

func GetStatsData() parser.Stats {
	return statsData
}

func GetEventLog() []parser.Event {
	return eventLog
}
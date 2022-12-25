package main

import (
	"log"
	"os"
	"strings"
)

func Import() (rucksacks []Rucksack) {
	rucksacks = make([]Rucksack, 0)
	inputBytes, err := os.ReadFile(inputPath)
	if err != nil {
		log.Panic(err)
	}
	for _, line := range strings.Split(strings.Replace(string(inputBytes), "\r", "", -1), "\n") {
		rucksacks = append(rucksacks, NewRucksack(line))
	}
	return
}

func ImportGroups() (groups [][]string) {
	groups = make([][]string, 0)
	inputBytes, err := os.ReadFile(inputPath)
	if err != nil {
		log.Panic(err)
	}
	var lines = strings.Split(strings.Replace(string(inputBytes), "\r", "", -1), "\n")
	for i := 0; i < len(lines); i += 3 {
		groups = append(groups, lines[i:i+3])
	}
	return
}

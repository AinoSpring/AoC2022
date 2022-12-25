package main

import "strings"

type Rucksack struct {
	compartment1 string
	compartment2 string
}

func NewRucksack(input string) Rucksack {
	return Rucksack{input[:(len(input) / 2)], input[(len(input) / 2):]}
}

func (rucksack *Rucksack) FindSame() (sameCharacters []string) {
	sameCharacters = make([]string, 0)
	for _, character := range strings.Split(rucksack.compartment1, "") {
		if strings.Contains(rucksack.compartment2, character) && (!strings.Contains(strings.Join(sameCharacters, ""), character)) {
			sameCharacters = append(sameCharacters, character)
		}
	}
	return
}

func FindAllSame(rucksacks []Rucksack) string {
	var sames = make([]string, 0)
	for _, rucksack := range rucksacks {
		sames = append(sames, rucksack.FindSame()...)
	}
	return strings.Join(sames, "")
}

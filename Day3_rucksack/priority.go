package main

import "strings"

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
)

var (
	priorities = append(strings.Split(alphabet, ""), strings.Split(strings.ToUpper(alphabet), "")...)
)

func Priority(character string) int {
	return SliceIndex(priorities, character) + 1
}

func PrioritySum(characterString string) (sum int) {
	for _, character := range strings.Split(characterString, "") {
		sum += Priority(character)
	}
	return
}

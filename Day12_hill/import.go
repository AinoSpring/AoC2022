package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Import() (elves [][]int) {
	elves = make([][]int, 0)
	contentBytes, err := os.ReadFile(input)
	if err != nil {
		log.Panic(err)
	}
	var content = strings.Replace(string(contentBytes), "\r", "", -1)
	var currentElve = make([]int, 0)
	for _, line := range strings.Split(content, "\n") {
		if line == "" {
			elves = append(elves, currentElve)
			currentElve = make([]int, 0)
			continue
		}
		calories, err := strconv.Atoi(line)
		if err != nil {
			log.Panic(err)
		}
		currentElve = append(currentElve, calories)
	}
	return
}

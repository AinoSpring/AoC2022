package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	directions = map[string]Vector{
		"U": {0, -1},
		"D": {0, 1},
		"L": {-1, 0},
		"R": {1, 0},
	}
)

func Import() (rope Rope) {
	rope = NewRope()
	contentBytes, err := os.ReadFile(input)
	if err != nil {
		log.Panic(err)
	}
	var content = strings.Replace(string(contentBytes), "\r", "", -1)
	for _, line := range strings.Split(content, "\n") {
		var command = strings.Split(line, " ")
		quantity, err := strconv.Atoi(command[1])
		if err != nil {
			log.Panic(err)
		}
		for i := 0; i < quantity; i++ {
			rope.Move(directions[command[0]])
			rope.UpdateTail()
		}
	}
	rope.RemoveTailPathDuplicates()
	return
}

func ImportLongerRope() (rope LongerRope) {
	rope = NewLongerRope(10)
	contentBytes, err := os.ReadFile(input)
	if err != nil {
		log.Panic(err)
	}
	var content = strings.Replace(string(contentBytes), "\r", "", -1)
	for _, line := range strings.Split(content, "\n") {
		var command = strings.Split(line, " ")
		quantity, err := strconv.Atoi(command[1])
		if err != nil {
			log.Panic(err)
		}
		for i := 0; i < quantity; i++ {
			rope.Move(directions[command[0]])
			rope.Update()
		}
	}
	rope.RemoveTailPathDuplicates()
	return
}

package main

import (
	"log"
	"os"
	"strings"
)

func Import() (games []Game) {
	games = make([]Game, 0)
	contentBytes, err := os.ReadFile(input)
	if err != nil {
		log.Panic(err)
	}
	var content = strings.Replace(string(contentBytes), "\r", "", -1)
	for _, line := range strings.Split(content, "\n") {
		if line == "" {
			continue
		}
		var game Game
		game.opponent = symbols[strings.Split(line, " ")[0]]
		game.player = symbols[strings.Split(line, " ")[1]]
		games = append(games, game)
	}
	return
}

func ImportUsingNewGame() (games []Game) {
	games = make([]Game, 0)
	contentBytes, err := os.ReadFile(input)
	if err != nil {
		log.Panic(err)
	}
	var content = strings.Replace(string(contentBytes), "\r", "", -1)
	for _, line := range strings.Split(content, "\n") {
		if line == "" {
			continue
		}
		games = append(games, NewGame(symbols[strings.Split(line, " ")[0]], strings.Split(line, " ")[1]))
	}
	return
}

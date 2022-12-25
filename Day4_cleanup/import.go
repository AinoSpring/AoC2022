package main

import (
	"log"
	"os"
	"strings"
)

func Import() (pairs []Pair) {
	pairs = make([]Pair, 0)
	contentBytes, err := os.ReadFile(input)
	if err != nil {
		log.Panic(err)
	}
	var content = strings.Replace(string(contentBytes), "\r", "", -1)
	for _, line := range strings.Split(content, "\n") {
		var a, b = strings.Split(line, ",")[0], strings.Split(line, ",")[1]
		pairs = append(pairs, Pair{NewCleaner(a), NewCleaner(b)})
	}
	return
}

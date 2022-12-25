package main

import (
	"log"
	"os"
	"strings"
)

func Import() []string {
	contentBytes, err := os.ReadFile(input)
	if err != nil {
		log.Panic(err)
	}
	var content = strings.Replace(string(contentBytes), "\r", "", -1)
	return strings.Split(strings.Replace(content, "\n", "", -1), "")
}

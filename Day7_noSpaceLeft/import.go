package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Import() (session Session) {
	session = NewSession()
	contentBytes, err := os.ReadFile(input)
	if err != nil {
		log.Panic(err)
	}
	var content = strings.Replace(string(contentBytes), "\r", "", -1)
	for _, line := range strings.Split(content, "\n") {
		if line == "" {
			continue
		}
		var command = strings.Split(line, " ")
		if command[0] == "$" {
			if command[1] == "cd" {
				session.currentWorkingDirectory.changeDirectory(command[2])
			}
			continue
		}
		if command[0] == "dir" {
			continue
		}
		size, err := strconv.Atoi(command[0])
		if err != nil {
			log.Panic(err)
		}
		session.NewFile(command[1], size)
	}
	return
}

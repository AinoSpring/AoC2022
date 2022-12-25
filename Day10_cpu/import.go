package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Import() (cpu CPU) {
	cpu = NewCPU()
	contentBytes, err := os.ReadFile(input)
	if err != nil {
		log.Panic(err)
	}
	var content = strings.Replace(string(contentBytes), "\r", "", -1)
	for _, line := range strings.Split(content, "\n") {
		var command = append(strings.Split(line, " "), "0")
		arg, err := strconv.Atoi(command[1])
		if err != nil {
			log.Panic(err)
		}
		cpu.Instruction(ParseInstruction(command[0]), arg)
	}
	return
}

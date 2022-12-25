package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Import() (instructions []string, stacks Stacks) {
	stacks = NewStacks()
	contentBytes, err := os.ReadFile(input)
	if err != nil {
		log.Panic(err)
	}
	var content = strings.Replace(string(contentBytes), "\r", "", -1)
	var lines = strings.Split(content, "\n")
	stacks.stacks = make([]Stack, (len(lines[0])+1)/4)
	for idx, line := range lines {
		if line == "" {
			instructions = lines[(idx + 1):]
			break
		}
		var isContainer = strings.Contains(line, "[")
		for i := 1; i < len(line); i += 4 {
			var n = (i - 1) / 4
			if isContainer {
				var container = strings.Split(line, "")[i]
				if container != " " {
					if stacks.stacks[n].crates == nil {
						stacks.stacks[n] = NewStack()
					}
					stacks.stacks[n].Append(container)
				}
				continue
			}
			label, err := strconv.Atoi(strings.Split(line, "")[i])
			if err != nil {
				log.Panic(err)
			}
			stacks.labels = append(stacks.labels, label)
		}
	}
	for i := 0; i < len(stacks.stacks); i++ {
		stacks.stacks[i].crates = SliceReverse(stacks.stacks[i].crates)
	}
	return
}

package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Import() (trees Trees) {
	trees = NewTrees()
	contentBytes, err := os.ReadFile(input)
	if err != nil {
		log.Panic(err)
	}
	var content = strings.Replace(string(contentBytes), "\r", "", -1)
	for _, line := range strings.Split(content, "\n") {
		for _, tree := range strings.Split(line, "") {
			treeValue, err := strconv.Atoi(tree)
			if err != nil {
				log.Panic(err)
			}
			trees.Append(treeValue)
		}
		trees.Row()
	}
	trees.RemoveLastRow()
	trees.FixSize()
	return
}

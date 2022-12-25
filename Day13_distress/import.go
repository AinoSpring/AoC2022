package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Import() (cTrees []CTree) {
	cTrees = make([]CTree, 0)
	contentBytes, err := os.ReadFile(input)
	if err != nil {
		log.Panic(err)
	}
	var content = string(contentBytes)
	var cCTreeIndex = NewCTreeIndex()
	var cCTree = NewCTree()
	for _, line := range strings.Split(strings.Replace(content+"\n", "\n\n", "\n", -1), "\n") {
		if line == "\r" {
			cTrees = append(cTrees, cCTree)
			cCTree = NewCTree()
			continue
		}
		line = strings.Replace(line, "\r", "", -1)
		for _, character := range strings.Split(line, "") {
			if character == "," {
				continue
			}
			if character == "[" {
				cCTreeIndex = append(cCTreeIndex, cCTree.Get(cCTreeIndex).New(NewCTree()))
			} else if character == "]" {
				cCTreeIndex.Pop()
			} else {
				intValue, err := strconv.Atoi(character)
				if err != nil {
					log.Panic(err)
				}
				cCTree.Get(cCTreeIndex).New(NewCTreeLeaf(intValue))
			}
		}
	}
	return
}

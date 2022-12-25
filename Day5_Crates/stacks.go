package main

import (
	"log"
	"strconv"
	"strings"
)

type Stacks struct {
	stacks []Stack
	labels []int
}

func NewStacks() (stacks Stacks) {
	stacks.stacks = make([]Stack, 0)
	stacks.labels = make([]int, 0)
	return
}

func (stacks *Stacks) Move(from int, to int) {
	stacks.stacks[SliceIndex(stacks.labels, to)].Append(stacks.stacks[SliceIndex(stacks.labels, from)].PopLast())
}

func (stacks *Stacks) MoveN(from int, to int, n int) {
	for i := 0; i < n; i++ {
		stacks.Move(from, to)
	}
}

func (stacks *Stacks) MoveNSameOrder(from int, to int, n int) {
	var moveCrates = make([]string, 0)
	for i := 0; i < n; i++ {
		moveCrates = append(moveCrates, stacks.stacks[SliceIndex(stacks.labels, from)].PopLast())
	}
	stacks.stacks[SliceIndex(stacks.labels, to)].crates = append(stacks.stacks[SliceIndex(stacks.labels, to)].crates, SliceReverse(moveCrates)...)
}

func (stacks *Stacks) ExecuteInstructions(instructions []string) {
	for _, instruction := range instructions {
		var command = strings.Split(instruction, " ")
		n, err := strconv.Atoi(command[1])
		if err != nil {
			log.Panic(err)
		}
		from, err := strconv.Atoi(command[3])
		if err != nil {
			log.Panic(err)
		}
		to, err := strconv.Atoi(command[5])
		if err != nil {
			log.Panic(err)
		}
		stacks.MoveN(from, to, n)
	}
}

func (stacks *Stacks) ExecuteInstructionsSameOrder(instructions []string) {
	for _, instruction := range instructions {
		var command = strings.Split(instruction, " ")
		n, err := strconv.Atoi(command[1])
		if err != nil {
			log.Panic(err)
		}
		from, err := strconv.Atoi(command[3])
		if err != nil {
			log.Panic(err)
		}
		to, err := strconv.Atoi(command[5])
		if err != nil {
			log.Panic(err)
		}
		stacks.MoveNSameOrder(from, to, n)
	}
}

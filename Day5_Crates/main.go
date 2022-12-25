package main

import (
	"fmt"
)

func main() {
	instructions, stacks := Import()
	stacks.ExecuteInstructions(instructions)
	fmt.Println()
	for _, stack := range stacks.stacks {
		fmt.Print(stack.PopLast())
	}
	fmt.Println()
	fmt.Print("\n------Part two------\n\n")
	instructions, stacks = Import()
	stacks.ExecuteInstructionsSameOrder(instructions)
	for _, stack := range stacks.stacks {
		fmt.Print(stack.PopLast())
	}
	fmt.Print("\n\n")
}

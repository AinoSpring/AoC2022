package main

import (
	"fmt"
)

func main() {
	var elves = Import()
	fmt.Println()
	fmt.Println(MaxCalories(Sum(elves)))
	fmt.Print("\n------Part two------\n\n")
	var total int
	for _, elve := range MaxN(Sum(elves), 3) {
		total += elve
	}
	fmt.Println(total)
	fmt.Println()
}

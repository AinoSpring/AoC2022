package main

import (
	"fmt"
)

func main() {
	var monkeys = Import()
	for i := 0; i < 20; i++ {
		monkeys.Update()
	}
	fmt.Println()
	fmt.Println(monkeys.MonkeyBusiness())
	fmt.Print("\n------Part two------\n\n")
	var newMonkeys = Import()
	newMonkeys.worryLevelDivision = false
	for i := 0; i < 10000; i++ {
		newMonkeys.Update()
	}
	fmt.Println(newMonkeys.MonkeyBusiness())
	fmt.Println()
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println(PrioritySum(FindAllSame(Import())))
	fmt.Println()
	fmt.Print("\n------Part two------\n\n")
	fmt.Println(PrioritySum(FindAllBadges(ImportGroups())))
	fmt.Println()
}

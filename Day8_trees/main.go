package main

import (
	"fmt"
)

func main() {
	var trees = Import()
	fmt.Println()
	fmt.Println(trees.CountVisible())
	fmt.Print("\n------Part two------\n\n")
	fmt.Println(trees.GetHighestScenicScore())
	fmt.Println()
}

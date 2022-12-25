package main

import "fmt"

func main() {
	fmt.Println()
	fmt.Println(CalculateScore(Import()))
	fmt.Print("\n------Part two------\n\n")
	fmt.Println(CalculateScore(ImportUsingNewGame()))
	fmt.Println()
}

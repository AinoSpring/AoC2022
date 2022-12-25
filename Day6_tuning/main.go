package main

import (
	"fmt"
)

func main() {
	var datastream = Import()
	fmt.Println()
	fmt.Println(FindMarker(datastream, 4))
	fmt.Print("\n------Part two------\n\n")
	fmt.Println(FindMarker(datastream, 14))
	fmt.Println()
}

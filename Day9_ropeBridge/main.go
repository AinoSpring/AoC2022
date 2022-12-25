package main

import (
	"fmt"
)

func main() {
	var rope = Import()
	fmt.Println()
	fmt.Println(rope.tailPath.Size()[1])
	fmt.Print("\n------Part two------\n\n")
	var longerRope = ImportLongerRope()
	fmt.Println(longerRope.tailPath.Size()[1])
	fmt.Println()
}

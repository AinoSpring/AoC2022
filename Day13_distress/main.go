package main

import (
	"fmt"
)

func main() {
	var cTrees = Import()
	fmt.Println()
	fmt.Println(IntSliceSum(CTreesRightIndices(cTrees)))
	fmt.Print("\n------Part two------\n\n")
}

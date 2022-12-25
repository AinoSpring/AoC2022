package main

import (
	"fmt"
)

func main() {
	var cpu = Import()
	fmt.Println()
	fmt.Println(cpu.XRegSum(xRegSumCycles))
	fmt.Print("\n------Part two------\n")
	fmt.Println(cpu.crt.display)
	fmt.Println()
}

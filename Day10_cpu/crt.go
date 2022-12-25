package main

import "math"

type CRT struct {
	display string
}

func (crt *CRT) Cycle(cycleN int, xReg int) {
	var displayX = (cycleN - 1) % 40
	if displayX == 0 {
		crt.display += "\n"
	}
	if math.Abs(float64(xReg-displayX)) <= 1 {
		crt.display += "#"
	} else {
		crt.display += "."
	}
}

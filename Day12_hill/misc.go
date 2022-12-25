package main

import (
	"math"
	"strings"
)

func SliceIndex[T comparable](slice []T, element T) int {
	for idx, cElement := range slice {
		if element == cElement {
			return idx
		}
	}
	return -1
}

func AlphabetIndex(character string) int {
	var alphabet = strings.Split("abcdefghijklmnopqrstuvwxyz", "")
	return SliceIndex(alphabet, character)
}

func Distance(aX int, aY int, bX int, bY int) int {
	return int(math.Abs(float64(aX-bX)) + math.Abs(float64(aY-bY)))
}

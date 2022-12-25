package main

func SliceIndex[T comparable](slice []T, element T) int {
	for idx, elem := range slice {
		if elem == element {
			return idx
		}
	}
	return -1
}

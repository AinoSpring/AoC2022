package main

func SliceCount[T comparable](slice []T, element T) (count int) {
	for _, elem := range slice {
		if elem == element {
			count++
		}
	}
	return
}

package main

func Remove[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func Pop(slice []int, index int) ([]int, int) {
	return Remove(slice, index), slice[index]
}

func Contains[T comparable](slice []T, element T) bool {
	for _, elem := range slice {
		if elem == element {
			return true
		}
	}
	return false
}

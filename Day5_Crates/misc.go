package main

func SliceIndex[T comparable](slice []T, element T) int {
	for idx, elem := range slice {
		if elem == element {
			return idx
		}
	}
	return -1
}

func SliceReverse[T any](s []T) []T {
	a := make([]T, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}

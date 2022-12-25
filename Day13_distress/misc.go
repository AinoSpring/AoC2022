package main

func IntSliceSum(slice []int) (sum int) {
	for _, element := range slice {
		sum += element
	}
	return sum
}

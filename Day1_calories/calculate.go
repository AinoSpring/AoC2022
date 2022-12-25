package main

func Sum(elves [][]int) (sum []int) {
	sum = make([]int, 0)
	for _, elve := range elves {
		var elveCalories = 0
		for _, calories := range elve {
			elveCalories += calories
		}
		sum = append(sum, elveCalories)
	}
	return
}

func MaxCalories(elves []int) (calories int) {
	for _, elve := range elves {
		if elve > calories {
			calories = elve
		}
	}
	return
}

func MaxN(elves []int, n int) (maxN []int) {
	maxN = make([]int, 0)
	for i := 0; i < n; i++ {
		var calories int
		for _, elve := range elves {
			if elve > calories && (!Contains(maxN, elve)) {
				calories = elve
			}
		}
		maxN = append(maxN, calories)
	}
	return
}

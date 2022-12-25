package main

type Pair struct {
	a Cleaner
	b Cleaner
}

func (pair *Pair) DoesOverlap() bool {
	return pair.a.Between(pair.b) || pair.b.Between(pair.a)
}

func CountContains(pairs []Pair) (count int) {
	for _, pair := range pairs {
		if pair.a.Contains(pair.b) || pair.b.Contains(pair.a) {
			count++
		}
	}
	return
}

func CountBetween(pairs []Pair) (count int) {
	for _, pair := range pairs {
		if pair.DoesOverlap() {
			count++
		}
	}
	return
}

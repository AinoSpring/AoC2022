package main

import (
	"sort"
)

type Monkeys struct {
	monkeys             []Monkey
	monkeyLabels        map[string]int
	worryLevelDivision  bool
	lowestCommonDivider uint64
}

func NewMonkeys() Monkeys {
	return Monkeys{make([]Monkey, 0), make(map[string]int, 0), true, 1}
}

func (monkeys *Monkeys) Append(label string, monkey Monkey) {
	monkeys.monkeyLabels[label] = len(monkeys.monkeys)
	monkeys.monkeys = append(monkeys.monkeys, monkey)
}

func (monkeys *Monkeys) Update() {
	for idx, monkey := range monkeys.monkeys {
		monkeys.monkeys[idx].inspectC += uint64(len(monkey.items))
		for itemIdx := range monkey.items {
			monkeys.monkeys[idx].items[itemIdx] = monkey.operation(monkey.items[itemIdx])
			if monkeys.worryLevelDivision {
				monkeys.monkeys[idx].items[itemIdx] /= uint64(3)
			} else {
				monkeys.monkeys[idx].items[itemIdx] %= monkeys.lowestCommonDivider
			}
			monkeys.monkeys[monkeys.monkeyLabels[monkey.throw[monkey.Test(itemIdx)]]].Catch(monkeys.monkeys[idx].items[itemIdx])
		}
		monkeys.monkeys[idx].items = make([]uint64, 0)
	}
}

func (monkeys *Monkeys) Last() *Monkey {
	return &monkeys.monkeys[len(monkeys.monkeys)-1]
}

func (monkeys *Monkeys) MonkeyBusiness() uint64 {
	var sorted = make([]Monkey, len(monkeys.monkeys))
	copy(sorted, monkeys.monkeys)
	sort.Slice(sorted, func(i int, j int) bool {
		return sorted[i].inspectC > sorted[j].inspectC
	})
	return sorted[0].inspectC * sorted[1].inspectC
}

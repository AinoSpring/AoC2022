package main

type Monkey struct {
	items     []uint64
	inspectC  uint64
	operation func(uint64) uint64
	test      uint64
	throw     map[bool]string
}

func NewMonkey() Monkey {
	return Monkey{items: make([]uint64, 0), throw: make(map[bool]string, 0)}
}

func (monkey *Monkey) Inspect(item int, worryLevelDivision bool) {
	monkey.items[item] = monkey.operation(monkey.items[item])
	if worryLevelDivision {
		monkey.items[item] /= 3
	}
	monkey.inspectC++
}

func (monkey *Monkey) Test(index int) bool {
	return monkey.items[index]%monkey.test == 0
}

func (monkey *Monkey) Catch(item uint64) {
	monkey.items = append(monkey.items, item)
}

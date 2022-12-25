package main

import "log"

type Trees struct {
	trees   [][]int
	visible [][]bool
	size    []int
}

func NewTrees() (trees Trees) {
	trees = Trees{trees: make([][]int, 0), visible: make([][]bool, 0), size: make([]int, 2)}
	trees.Row()
	return
}

func (trees *Trees) At(position []int) int {
	return trees.trees[position[0]][position[1]]
}

func (trees *Trees) Visible(position []int) {
	trees.visible[position[0]][position[1]] = true
}

func (trees *Trees) Append(value int) {
	trees.trees[len(trees.trees)-1] = append(trees.trees[len(trees.trees)-1], value)
	trees.visible[len(trees.visible)-1] = append(trees.visible[len(trees.visible)-1], false)
}

func (trees *Trees) Row() {
	trees.trees = append(trees.trees, make([]int, 0))
	trees.visible = append(trees.visible, make([]bool, 0))
}

func (trees *Trees) RemoveLastRow() {
	trees.trees = trees.trees[:(len(trees.trees) - 1)]
	trees.visible = trees.visible[:(len(trees.visible) - 1)]
}

func (trees *Trees) FixSize() {
	trees.size = []int{len(trees.trees), len(trees.trees[0])}
}

func (trees *Trees) CountVisibleOnSideRow(side int, row int) {
	var delta []int
	var currentPosition []int
	var size int
	switch side {
	case 0:
		delta = []int{0, 1}
		currentPosition = []int{row, 0}
		size = trees.size[1]
	case 1:
		delta = []int{1, 0}
		currentPosition = []int{0, row}
		size = trees.size[0]
	case 2:
		delta = []int{0, -1}
		currentPosition = []int{row, trees.size[1] - 1}
		size = trees.size[1]
	case 3:
		delta = []int{-1, 0}
		currentPosition = []int{trees.size[0] - 1, row}
		size = trees.size[0]
	default:
		log.Panicf("Invalid side: %v", side)
		return
	}
	var oldValue = -1
	for i := 0; i < size; i++ {
		var currentValue = trees.At(currentPosition)
		if currentValue > oldValue {
			oldValue = currentValue
			trees.Visible(currentPosition)
		}
		currentPosition = []int{currentPosition[0] + delta[0], currentPosition[1] + delta[1]}
	}
}

func (trees *Trees) CountVisibleOnSide(side int) {
	var size int
	if side%2 == 0 {
		size = trees.size[0]
	} else {
		size = trees.size[1]
	}
	for i := 0; i < size; i++ {
		trees.CountVisibleOnSideRow(side, i)
	}
}

func (trees *Trees) CountVisible() (count int) {
	for i := 0; i < 4; i++ {
		trees.CountVisibleOnSide(i)
	}
	for x := 0; x < len(trees.visible); x++ {
		for y := 0; y < len(trees.visible); y++ {
			if trees.visible[x][y] {
				count++
			}
		}
	}
	return
}

func (trees *Trees) GetLengthInDirection(position []int, direction []int) (length int) {
	var height = trees.At(position)
	position = []int{position[0] + direction[0], position[1] + direction[1]}
	for position[0] < trees.size[0] && position[1] < trees.size[1] && position[0] >= 0 && position[1] >= 0 {
		var currentHeight = trees.At(position)
		length++
		if currentHeight >= height {
			break
		}
		position = []int{position[0] + direction[0], position[1] + direction[1]}
	}
	return
}

func (trees *Trees) GetLengths(position []int) (lengths []int) {
	var directions = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	lengths = make([]int, 0)
	for i := 0; i < 4; i++ {
		var tempPosition = make([]int, 2)
		copy(tempPosition, position)
		lengths = append(lengths, trees.GetLengthInDirection(tempPosition, directions[i]))
	}
	return
}

func (trees *Trees) GetHighestScenicScore() (max int) {
	var scenicScores = make([]int, 0)
	for x := 0; x < trees.size[0]; x++ {
		for y := 0; y < trees.size[1]; y++ {
			var scenicScore = 1
			for _, length := range trees.GetLengths([]int{x, y}) {
				scenicScore *= length
			}
			scenicScores = append(scenicScores, scenicScore)
		}
	}
	for _, scenicScore := range scenicScores {
		if scenicScore > max {
			max = scenicScore
		}
	}
	return
}

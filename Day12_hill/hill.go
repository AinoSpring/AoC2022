package main

import "sort"

type Hill struct {
	grid       [][]string
	positionX  int
	positionY  int
	bPositionX int
	bPositionY int
}

func NewHill() Hill {
	return Hill{grid: make([][]string, 0)}
}

func (hill *Hill) Get(x int, y int) string {
	return hill.grid[y][x]
}

func (hill *Hill) CheckPosition(positionX int, positionY int) bool {
	return positionX >= 0 && positionX < len(hill.grid[0]) && positionY >= 0 && positionY < len(hill.grid)
}

func (hill *Hill) AdjacentPositions() (aPositions [][]int) {
	aPositions = make([][]int, 0)
	for x := -1; x < 2; x += 2 {
		var cPositionX = hill.positionX + x
		if hill.CheckPosition(cPositionX, hill.positionY) {
			aPositions = append(aPositions, []int{cPositionX, hill.positionX})
		}
	}
	for y := -1; y < 2; y += 2 {
		var cPositionY = hill.positionY + y
		if hill.CheckPosition(hill.positionX, cPositionY) {
			aPositions = append(aPositions, []int{hill.positionX, cPositionY})
		}
	}
	return
}

func (hill *Hill) MoveUp() {
	var adjacentPositions = hill.AdjacentPositions()
	sort.Slice(adjacentPositions, func(i int, j int) bool {
		return Distance(adjacentPositions[i][0], adjacentPositions[i][1], hill.bPositionX, hill.bPositionY) < Distance(adjacentPositions[j][0], adjacentPositions[j][1], hill.bPositionX, hill.bPositionY)
	})
	var bestX, bestY = -1, -1
	for _, aPosition := range adjacentPositions {
		if bestX < 0 || bestY < 0 {
			bestX = aPosition[0]
			bestY = aPosition[1]
			continue
		}
		if hill.Get(bestX, bestY) < hill.Get(aPosition[0], aPosition[1]) {
			bestX = aPosition[0]
			bestY = aPosition[1]
		}
	}
	hill.positionX = bestX
	hill.positionY = bestY
}

package main

import (
	"log"
	"strconv"
)

type CTreeIndex []uint

type CTree struct {
	value int
	nodes []CTree
}

func NewCTreeIndex() CTreeIndex {
	return make(CTreeIndex, 0)
}

func (index *CTreeIndex) Pop() (value uint) {
	value = (*index)[len(*index)-1]
	*index = (*index)[:(len(*index) - 1)]
	return
}

func NewCTree() CTree {
	return CTree{nodes: make([]CTree, 0), value: -1}
}

func NewCTreeLeaf(value int) CTree {
	return CTree{nodes: make([]CTree, 0), value: value}
}

func (cTree *CTree) Get(index CTreeIndex) *CTree {
	if len(index) == 0 {
		return cTree
	}
	return cTree.nodes[index[0]].Get(index[1:])
}

func (cTree *CTree) New(tree CTree) uint {
	cTree.nodes = append(cTree.nodes, tree)
	return uint(len(cTree.nodes) - 1)
}

func (cTree *CTree) RightOrder() bool {
	return Compare(cTree.nodes[0], cTree.nodes[1]) != 0
}

func Compare(left CTree, right CTree) uint8 { // 0 FALSE ; 1 TRUE ; 2 SAME
	log.Printf("CMP: %v & %v", left.String(), right.String())
	for i := 0; ; i++ {
		if i >= len(left.nodes) || i >= len(right.nodes) {
			if len(left.nodes) < len(right.nodes) {
				return 1
			} else if len(left.nodes) > len(right.nodes) {
				return 0
			}
			return 2
		}
		var cLeft = left.nodes[i]
		var cRight = right.nodes[i]
		if cLeft.value != -1 && cRight.value != -1 {
			if cLeft.value != cRight.value {
				if cLeft.value < cRight.value {
					return 1
				} else {
					return 0
				}
			}
		}
		if cLeft.value != -1 && cRight.value == -1 {
			if result := Compare(CTree{
				value: -1,
				nodes: []CTree{NewCTreeLeaf(cLeft.value)},
			}, cRight); result != 2 {
				return result
			}
		}
		if cLeft.value == -1 && cRight.value != -1 {
			if result := Compare(cLeft, CTree{
				value: -1,
				nodes: []CTree{NewCTreeLeaf(cRight.value)},
			}); result != 2 {
				return result
			}
		}
		if cLeft.value == -1 && cRight.value == -1 {
			if result := Compare(cLeft, cRight); result != 2 {
				return result
			}
		}
	}
	return 2
}

func (cTree *CTree) String() string {
	var result = "["
	for _, node := range cTree.nodes {
		if node.value == -1 {
			result += node.String()
		} else {
			result += strconv.Itoa(node.value)
		}
	}
	return result + "]"
}

func CTreesRightIndices(cTrees []CTree) (indices []int) {
	indices = make([]int, 0)
	for idx, cTree := range cTrees {
		if cTree.RightOrder() {
			indices = append(indices, idx+1)
		}
	}
	return
}

package main

type Rope struct {
	head     Vector
	tail     Vector
	tailPath Matrix
}

func NewRope() Rope {
	return Rope{head: NewVector(2), tail: NewVector(2), tailPath: Matrix{NewVector(2)}}
}

func (rope *Rope) Move(delta Vector) {
	rope.head = rope.head.Add(delta)
}

func (rope *Rope) TailNextToHead() bool {
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if rope.head.Equal(rope.tail.Add(Vector{float64(x), float64(y)})) {
				return true
			}
		}
	}
	return false
}

func (rope *Rope) UpdateTail() {
	for !rope.TailNextToHead() {
		if rope.head[0] > rope.tail[0] {
			rope.tail[0] += 1
		} else if rope.head[0] < rope.tail[0] {
			rope.tail[0] -= 1
		}
		if rope.head[1] > rope.tail[1] {
			rope.tail[1] += 1
		} else if rope.head[1] < rope.tail[1] {
			rope.tail[1] -= 1
		}
		rope.tailPath = append(rope.tailPath, rope.tail.Copy())
	}
}

func (rope *Rope) RemoveTailPathDuplicates() {
	var newTailPath = NewMatrix(NewVector(2))
	for _, vector := range rope.tailPath {
		if newTailPath.Count(vector) < 1 {
			newTailPath = append(newTailPath, vector)
		}
	}
	rope.tailPath = newTailPath
}

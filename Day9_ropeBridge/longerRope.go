package main

type LongerRope struct {
	rope     Matrix
	tailPath Matrix
	length   int
}

func NewLongerRope(length int) LongerRope {
	return LongerRope{rope: NewMatrix(Vector{2, float64(length)}), tailPath: Matrix{NewVector(2)}, length: length}
}

func (rope *LongerRope) Move(delta Vector) {
	rope.rope[0] = rope.rope[0].Add(delta)
}

func (rope *LongerRope) NextToNext(a int) bool {
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if rope.rope[a-1].Equal(rope.rope[a].Add(Vector{float64(x), float64(y)})) {
				return true
			}
		}
	}
	return false
}

func (rope *LongerRope) UpdateSegment(segment int) {
	for !rope.NextToNext(segment) {
		var next = rope.rope[segment-1]
		if next[0] > rope.rope[segment][0] {
			rope.rope[segment][0] += 1
		} else if next[0] < rope.rope[segment][0] {
			rope.rope[segment][0] -= 1
		}
		if next[1] > rope.rope[segment][1] {
			rope.rope[segment][1] += 1
		} else if next[1] < rope.rope[segment][1] {
			rope.rope[segment][1] -= 1
		}
		if segment == rope.length-1 {
			rope.tailPath = append(rope.tailPath, rope.rope[segment].Copy())
		}
	}
}

func (rope *LongerRope) Update() {
	for i := 1; i < rope.length; i++ {
		rope.UpdateSegment(i)
	}
}

func (rope *LongerRope) RemoveTailPathDuplicates() {
	var newTailPath = NewMatrix(NewVector(2))
	for _, vector := range rope.tailPath {
		if newTailPath.Count(vector) < 1 {
			newTailPath = append(newTailPath, vector)
		}
	}
	rope.tailPath = newTailPath
}

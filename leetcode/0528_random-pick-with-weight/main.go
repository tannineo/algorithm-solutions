package main

import (
	"math/rand"
	"time"
)

type Solution struct {
	r        *rand.Rand
	rangeArr []int
}

func Constructor(w []int) Solution {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	lenR := len(w)
	rangeArr := make([]int, lenR)
	copy(rangeArr, w)
	for i := 1; i < lenR; i++ {
		rangeArr[i] = rangeArr[i-1] + rangeArr[i]
	}
	return Solution{rand, rangeArr}
}

func (this *Solution) PickIndex() int {
	pick := this.r.Float64() * float64(this.rangeArr[len(this.rangeArr)-1])

	if pick < float64(this.rangeArr[0]) {
		return 0
	}

	left, right := 1, len(this.rangeArr)
	var guess int
	for left < right {
		guess = (left + right) / 2
		if float64(this.rangeArr[guess-1]) <= pick && pick < float64(this.rangeArr[guess]) {
			return guess
		}
		if pick < float64(this.rangeArr[guess-1]) {
			right = guess
		}
		if pick >= float64(this.rangeArr[guess]) {
			left = guess + 1
		}
	}

	return guess
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */

func main() {
	solu := Constructor([]int{1})
	solu.PickIndex()
}

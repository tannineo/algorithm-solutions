package main

import (
	rand "math/rand"
	"time"
)

type RandomizedSet struct {
	indexMap map[int]int
	dataArr  []int
	lenD     int
	rand     *rand.Rand
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return RandomizedSet{map[int]int{}, []int{}, 0, rand}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.indexMap[val]; !ok {
		this.dataArr = append(this.dataArr, val)
		this.indexMap[val] = this.lenD
		this.lenD++
		return true
	}
	return false
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if idx, ok := this.indexMap[val]; ok {
		this.dataArr[idx] = this.dataArr[this.lenD-1]
		this.lenD--
		this.dataArr = this.dataArr[:this.lenD]
		delete(this.indexMap, val)
		if idx < this.lenD { // if not last element
			this.indexMap[this.dataArr[idx]] = idx
		}
		return true
	}
	return false
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	// it's guaranteed that
	// at least one element exists when this method is called
	if this.lenD == 0 {
		return 0
	}

	return this.dataArr[this.rand.Intn(this.lenD)]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {
	r := Constructor()

	println(r.Insert(1)) // true

	println(r.Remove(2)) // false

	println(r.Insert(2)) // true

	println(r.GetRandom()) // 1 or 2

	println(r.Remove(1)) // true

	println(r.Insert(2)) // false

	println(r.GetRandom()) // 2
}

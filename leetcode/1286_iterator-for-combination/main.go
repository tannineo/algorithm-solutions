package main

type CombinationIterator struct {
	NextArr     []int
	LastArr     []int
	HasNextBool bool
	Runes       []rune
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	NextArr := make([]int, combinationLength)
	LastArr := make([]int, combinationLength)
	Runes := []rune(characters)
	for i := range NextArr {
		NextArr[i] = i
	}
	cur := len(Runes) - len(NextArr)
	for i := range LastArr {
		LastArr[i] = cur
		cur++
	}
	return CombinationIterator{NextArr, LastArr, true, Runes}
}

func (this *CombinationIterator) Next() string {
	result := []rune{}

	for _, num := range this.NextArr {
		result = append(result, this.Runes[num])
	}

	// check if is the last
	this.HasNextBool = false
	for i := range this.NextArr {
		if this.NextArr[i] != this.LastArr[i] {
			this.HasNextBool = true
		}
	}

	if this.HasNext() {
		this.getNextArr()
	}

	return string(result)
}

func (this *CombinationIterator) getNextArr() {
	upper := len(this.Runes) - 1
	for {
		carry := 1
		cur := len(this.NextArr) - 1
		for carry > 0 && cur >= 0 {
			this.NextArr[cur] += carry
			carry = 0
			if this.NextArr[cur] > upper {
				carry = 1
				this.NextArr[cur] = 0
			}
			cur--
		}
		// check if all the numbers are different and in increasing order
		flag := false
		for i := 1; i < len(this.NextArr); i++ {
			if this.NextArr[i-1] >= this.NextArr[i] {
				flag = true
				break
			}
		}
		if !flag { // all increasing
			break
		}
	}
}

func (this *CombinationIterator) HasNext() bool {
	return this.HasNextBool
}

/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
func main() {
	obj := Constructor("abc", 2)
	println(obj.Next())    // ab
	println(obj.HasNext()) // true
	println(obj.Next())    // ac
	println(obj.HasNext()) // true
	println(obj.Next())    // bc
	println(obj.HasNext()) // false
}

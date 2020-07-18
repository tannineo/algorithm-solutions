package main

type Element struct {
	Val    int
	MinVal int
}

type MinStack struct {
	arr []Element
	top int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{[]Element{}, -1}
}

func (this *MinStack) Push(x int) {
	if this.top < 0 {
		// the stack is empty to add element
		this.arr = append(this.arr, Element{Val: x, MinVal: x})
	} else {
		// the stack is not empty
		this.arr = append(this.arr, Element{Val: x, MinVal: min(x, this.arr[this.top].MinVal)})
	}
	this.top++
}

func (this *MinStack) Pop() {
	this.arr = this.arr[:this.top]
	this.top--
}

func (this *MinStack) Top() int {
	return this.arr[this.top].Val
}

func (this *MinStack) GetMin() int {
	return this.arr[this.top].MinVal
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)

	obj.Push(3)

	obj.Push(-1)

	obj.Push(-2)

	println(obj.Top(), "", obj.GetMin())
	obj.Pop()

	println(obj.Top(), "", obj.GetMin())

	obj.Push(-3)
	println(obj.Top(), "", obj.GetMin())

	obj.Pop()

	println(obj.Top(), "", obj.GetMin())

	obj.Push(1)
	println(obj.Top(), "", obj.GetMin())

	obj.Push(2)

	println(obj.Top(), "", obj.GetMin())

}

package main

type Node struct {
	Val int
	Pos int

	Prev *Node
	Next *Node
}

type FirstUnique struct {
	orderMap map[int]*Node
	existMap map[int]struct{}
	head     *Node // init -1
	tail     *Node // init -1
	count    int
}

func Constructor(nums []int) FirstUnique {
	orderMap := make(map[int]*Node)
	existMap := make(map[int]struct{})

	headNode := Node{-1, -1, nil, nil}
	tailNode := Node{-1, -1, nil, nil}
	tailNode.Prev = &headNode
	headNode.Next = &tailNode

	firstUnique := FirstUnique{orderMap, existMap, &headNode, &tailNode, 0}

	for _, num := range nums {
		firstUnique.Add(num)
	}

	return firstUnique
}

func (this *FirstUnique) ShowFirstUnique() int {
	return this.head.Next.Val
}

func (this *FirstUnique) Add(value int) {
	// 0. incr count
	this.count++

	if _, ok := this.existMap[value]; ok {
		// another value

		if firstUniqueNode, ok := this.orderMap[value]; ok {
			// was unique before

			// 1. pick from the linkedList
			firstUniqueNode.Prev.Next = firstUniqueNode.Next
			firstUniqueNode.Next.Prev = firstUniqueNode.Prev

			// 2. delete from the orderMap
			delete(this.orderMap, value)
		}
		// not unique before
		// do nothing

	} else {
		// new value

		// 1. record existence
		this.existMap[value] = struct{}{}
		// 2. create a new Node and append it to tail
		newNode := Node{value, this.count - 1, this.tail.Prev, this.tail}
		this.tail.Prev.Next = &newNode
		this.tail.Prev = &newNode
		// 3. add it into the orderMap
		this.orderMap[value] = &newNode
	}
}

func main() {
	firstUnique := Constructor([]int{7, 7, 7, 7, 7, 7})
	println(firstUnique.ShowFirstUnique()) // return -1
	firstUnique.Add(7)                     // the queue is now [7,7,7,7,7,7,7]
	firstUnique.Add(3)                     // the queue is now [7,7,7,7,7,7,7,3]
	firstUnique.Add(3)                     // the queue is now [7,7,7,7,7,7,7,3,3]
	firstUnique.Add(7)                     // the queue is now [7,7,7,7,7,7,7,3,3,7]
	firstUnique.Add(17)                    // the queue is now [7,7,7,7,7,7,7,3,3,7,17]
	println(firstUnique.ShowFirstUnique()) // return 17
}

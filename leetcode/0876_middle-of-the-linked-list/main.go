package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	fastWalker := head
	slowWalker := head

	count := 0

	for fastWalker.Next != nil {
		fastWalker = fastWalker.Next
		count++

		if count%2 == 1 {
			slowWalker = slowWalker.Next
		}
	}

	return slowWalker
}

func main() {
}

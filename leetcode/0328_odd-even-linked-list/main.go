package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var oddFirst, oddTail *ListNode
	var evenFirst, evenTail *ListNode

	// init
	oddFirst = head
	oddTail = head
	counter := 1
	curNode := head.Next

	for curNode != nil {
		counter++
		if counter&1 == 0 { // even
			if evenFirst == nil {
				evenFirst = curNode
			} else {
				evenTail.Next = curNode
			}
			evenTail = curNode
		} else { // odd
			oddTail.Next = curNode
			oddTail = curNode
		}

		curNode = curNode.Next
	}

	// combine
	oddTail.Next = evenFirst
	evenTail.Next = nil

	return oddFirst
}

func main() {}

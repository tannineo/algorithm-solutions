package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	a := head
	b := head.Next
	a.Next = b.Next
	b.Next = a
	head = b

	prev := a
	if prev.Next == nil {
		return head
	}
	a = prev.Next
	b = a.Next

	for b != nil {

		prev.Next = b
		a.Next = b.Next
		b.Next = a

		prev = a
		if prev.Next == nil {
			return head
		}
		a = prev.Next
		b = a.Next
	}

	return head
}

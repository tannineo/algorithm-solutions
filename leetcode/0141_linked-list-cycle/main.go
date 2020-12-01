package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	a, b := head, head.Next

	for b != nil && a != b {
		a = a.Next
		b = b.Next
		if b != nil {
			b = b.Next
		} else { // b comes to the end
			return false
		}
	}

	return b != nil
}

package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	a, b := head, head

	// phase 1
	for a != nil && b != nil {
		a = a.Next
		b = b.Next
		if b == nil {
			return nil
		}
		b = b.Next
		if a == nil || b == nil {
			return nil
		}

		if a == b {
			break
		}
	}

	// loop exists, phase 2

	a = head
	for a != b {
		a = a.Next
		b = b.Next
	}

	return a
}

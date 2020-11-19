package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head.Next
	head.Next = nil

	for cur != nil {
		last2, last1 := (*ListNode)(nil), head
		for last1 != cur && last1 != nil && last1.Val < cur.Val {
			last2, last1 = last1, last1.Next
		}

		if last2 == nil {
			head = cur
			cur = cur.Next
			head.Next = last1
			continue
		}

		last2.Next = cur
		cur = cur.Next
		last2.Next.Next = last1
	}

	return head
}

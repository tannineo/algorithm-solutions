package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	cur1, cur2 := l1, l2

	var newList *ListNode
	if cur1.Val < cur2.Val {
		newList = cur1
		cur1 = cur1.Next
	} else {
		newList = cur2
		cur2 = cur2.Next
	}

	curNew := newList

	for cur1 != nil && cur2 != nil {
		if cur1.Val < cur2.Val {
			curNew.Next = cur1
			cur1 = cur1.Next
		} else {
			curNew.Next = cur2
			cur2 = cur2.Next
		}
		curNew = curNew.Next
	}

	// pin the rest
	if cur1 != nil {
		curNew.Next = cur1
	} else if cur2 != nil {
		curNew.Next = cur2
	}

	return newList
}

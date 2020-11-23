package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// divide by fast/slow pointers
	fast, slow := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil // to break the link and divide
	return mergeSortedList(sortList(slow), sortList(mid))
}

func mergeSortedList(head1, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	} else if head2 == nil {
		return head1
	}

	var newList, newCur *ListNode
	if head1.Val < head2.Val {
		newList = head1
		head1 = head1.Next
		newList.Next = nil
		newCur = newList
	} else {
		newList = head2
		head2 = head2.Next
		newList.Next = nil
		newCur = newList
	}

	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			newCur.Next = head1
			head1 = head1.Next
			newCur = newCur.Next
			newCur.Next = nil
		} else {
			newCur.Next = head2
			head2 = head2.Next
			newCur = newCur.Next
			newCur.Next = nil
		}
	}

	return newList
}

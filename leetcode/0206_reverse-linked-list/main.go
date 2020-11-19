package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	newHead, _ := reverseListRecursive(head)
	// newHead = reverseListIterative(head)
	return newHead
}

func reverseListRecursive(head *ListNode) (newHead, newTail *ListNode) {
	if head == nil {
		return nil, nil
	}

	if head.Next == nil {
		return head, head
	}

	newHead, newTail = reverseListRecursive(head.Next)
	newTail.Next = head
	head.Next = nil

	return newHead, head
}

func reverseListIterative(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	listNodes := []*ListNode{head}

	for head.Next != nil {
		listNodes = append(listNodes, head.Next)
		head = head.Next
	}

	newCur := listNodes[len(listNodes)-1]
	newCur.Next = nil
	newHead := newCur
	for i := len(listNodes) - 2; i >= 0; i-- {
		newCur.Next = listNodes[i]
		newCur = newCur.Next
		newCur.Next = nil
	}

	return newHead
}

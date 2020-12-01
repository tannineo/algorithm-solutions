package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// nil list?
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	len1, len2 := listLen(l1), listLen(l2)

	if len2 > len1 { // swap
		l1, l2 = l2, l1
		len1, len2 = len2, len1
	}

	// sum Vals to a new list
	var newList, curNodeNew *ListNode

	cur := len1
	curNode1 := l1
	for cur > len2 {
		if newList == nil {
			newList = &ListNode{}
			curNodeNew = newList
		} else {
			curNodeNew.Next = &ListNode{}
			curNodeNew = curNodeNew.Next
		}

		curNodeNew.Val = curNode1.Val
		cur--
		curNode1 = curNode1.Next
	}

	curNode2 := l2
	for cur > 0 {
		if newList == nil { // for len1 == len2
			newList = &ListNode{}
			curNodeNew = newList
		} else {
			curNodeNew.Next = &ListNode{}
			curNodeNew = curNodeNew.Next
		}

		curNodeNew.Val = curNode1.Val + curNode2.Val
		curNode1 = curNode1.Next
		curNode2 = curNode2.Next
		cur--
	}

	// deal with the carry, recursively
	carry := carryForList(newList)
	if carry == 1 {
		newList = &ListNode{1, newList}
	}

	return newList
}

func listLen(l *ListNode) int {
	ans := 0
	for l != nil {
		ans++
		l = l.Next
	}
	return ans
}

func carryForList(l *ListNode) int { // return 1 or 0
	if l == nil {
		return 0
	}

	c := carryForList(l.Next)

	l.Val = l.Val + c
	if l.Val >= 10 {
		l.Val -= 10
		return 1
	}
	return 0
}

func array2List(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	return &ListNode{arr[0], array2List(arr[1:])}
}

func main() {
	addTwoNumbers(array2List([]int{7, 2, 4, 3}), array2List([]int{5, 6, 4}))
}

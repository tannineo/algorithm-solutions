package main

import "container/heap"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Trying to implement priority queue
type ListNodePQ []*ListNode

// sort.Interface
func (pq ListNodePQ) Len() int { return len(pq) }

func (pq ListNodePQ) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq ListNodePQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// heap.Interface
func (pq *ListNodePQ) Push(x interface{}) {
	item := x.(*ListNode)
	*pq = append(*pq, item)
}

func (pq *ListNodePQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[:n-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {

	pq := (ListNodePQ)([]*ListNode{})
	var newList, newCur *ListNode

	heap.Init(&pq)

	// push
	for _, v := range lists {
		if v != nil {
			heap.Push(&pq, v)
		}
	}

	// loop pop
	for len(pq) > 0 {
		curHead := heap.Pop(&pq).(*ListNode)
		if curHead.Next != nil {
			heap.Push(&pq, curHead.Next)
		}
		if newList == nil {
			newList = curHead
			newCur = newList
		} else {
			newCur.Next = curHead
			newCur = newCur.Next
			newCur.Next = nil
		}
	}

	return newList
}

package main

import (
	"fmt"
	"strconv"
)

type MyLinkedList struct {
	head, tail *Node
	len        int
}

type Node struct {
	Val        int
	Prev, Next *Node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.len { // invalid index
		return -1
	}

	var cur int
	var curNode *Node
	if index < this.len/2 {
		// start from head
		cur = 0
		curNode = this.head
		for cur != index {
			curNode = curNode.Next
			cur++
		}
	} else {
		// start from tail
		cur = this.len - 1
		curNode = this.tail
		for cur != index {
			curNode = curNode.Prev
			cur--
		}
	}

	return curNode.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	if this.len == 0 {
		this.head = &Node{val, nil, nil}
		this.tail = this.head
	} else {
		this.head.Prev = &Node{Val: val, Prev: nil, Next: this.head}
		this.head = this.head.Prev
	}
	this.len++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this.len == 0 {
		this.head = &Node{val, nil, nil}
		this.tail = this.head
	} else {
		this.tail.Next = &Node{Val: val, Prev: this.tail, Next: nil}
		this.tail = this.tail.Next
	}
	this.len++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.len { // invalid index
		return
	}

	if index == 0 { // including initializing with 1st element inserted
		this.AddAtHead(val)
		return
	}

	if index == this.len { // including initializing with 1st element inserted
		this.AddAtTail(val)
		return
	}

	var cur int
	var curNode *Node
	if index < this.len/2 {
		// start from head
		cur = 0
		curNode = this.head
		for cur != index {
			curNode = curNode.Next
			cur++
		}
	} else {
		// start from tail
		cur = this.len - 1
		curNode = this.tail
		for cur != index {
			curNode = curNode.Prev
			cur--
		}
	}

	curNode.Prev.Next = &Node{Val: val, Prev: curNode.Prev, Next: curNode}
	curNode.Prev = curNode.Prev.Next
	this.len++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.len { // invalid index
		return
	}

	if this.len == 1 { // empty after deletion
		this.len = 0
		this.head = nil
		this.tail = nil
		return
	}

	switch index {
	case 0: // head
		this.head = this.head.Next
		this.head.Prev = nil

	case this.len - 1: // tail
		this.tail = this.tail.Prev
		this.tail.Next = nil

	default:
		var cur int
		var curNode *Node
		if index < this.len/2 {
			// start from head
			cur = 0
			curNode = this.head
			for cur != index {
				curNode = curNode.Next
				cur++
			}
		} else {
			// start from tail
			cur = this.len - 1
			curNode = this.tail
			for cur != index {
				curNode = curNode.Prev
				cur--
			}
		}
		curNode.Prev.Next = curNode.Next
		curNode.Next.Prev = curNode.Prev
	}
	this.len--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func (this MyLinkedList) String() string {
	str := fmt.Sprintf("len %v,", this.len)
	str += " ["
	cur := this.head
	for cur != nil {
		str += strconv.Itoa(cur.Val)
		cur = cur.Next
		if cur != nil {
			str += ", "
		}
	}
	str += "]"

	return str
}

func main() {
	// obj := Constructor()
	// fmt.Printf("%v\n", obj)
	// obj.AddAtHead(1)
	// fmt.Printf("%v\n", obj)
	// obj.AddAtTail(3)
	// fmt.Printf("%v\n", obj)
	// obj.AddAtIndex(1, 2)
	// fmt.Printf("%v\n", obj)
	// obj.DeleteAtIndex(1)
	// fmt.Printf("%v\n", obj)

	obj := Constructor()
	fmt.Printf("%v\n", obj)
	obj.AddAtHead(7)
	fmt.Printf("%v\n", obj)
	obj.AddAtHead(2)
	fmt.Printf("%v\n", obj)
	obj.AddAtHead(1)
	fmt.Printf("%v\n", obj)

	obj.AddAtIndex(3, 0)
	fmt.Printf("%v\n", obj)
	obj.DeleteAtIndex(2)
	fmt.Printf("%v\n", obj)

	obj.AddAtHead(6)
	fmt.Printf("%v\n", obj)

	obj.AddAtTail(4)
	fmt.Printf("%v\n", obj)
	fmt.Printf("%v\n", obj.Get(4))
	obj.AddAtHead(4)
	fmt.Printf("%v\n", obj)

	obj.AddAtIndex(5, 0)
	fmt.Printf("%v\n", obj)
	obj.AddAtHead(6)
	fmt.Printf("%v\n", obj)
}

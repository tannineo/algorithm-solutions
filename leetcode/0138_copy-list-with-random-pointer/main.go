package main

// Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {

	if head == nil {
		return nil
	}

	// create new nodes with Val
	// use a map to record: Val -> *Node
	mapVal2Node := map[*Node]*Node{}
	cur := head
	for cur != nil {
		mapVal2Node[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}

	// iterate the original list again to copy Next and Random
	cur = head
	for cur != nil {
		if cur.Next != nil {
			mapVal2Node[cur].Next = mapVal2Node[cur.Next]
		}
		if cur.Random != nil {
			mapVal2Node[cur].Random = mapVal2Node[cur.Random]
		}
		cur = cur.Next
	}

	return mapVal2Node[head]
}

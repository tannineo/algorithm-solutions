package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}

	str := strconv.Itoa(root.Val)

	str += " " + this.serialize(root.Left)
	str += " " + this.serialize(root.Right)

	return str
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	strS := strings.Split(data, " ")
	if strS[0] == "#" {
		return nil
	}
	val, err := strconv.Atoi(strS[0])
	if err != nil {
		return nil
	}
	newNode := TreeNode{val, nil, nil}

	cur := 1
	newNode.Left = this.deserializeRecur(strS, &cur)
	newNode.Right = this.deserializeRecur(strS, &cur)

	return &newNode
}

func (this *Codec) deserializeRecur(strS []string, cur *int) *TreeNode {
	if strS[*cur] == "#" {
		*cur++
		return nil
	}
	val, err := strconv.Atoi(strS[*cur])
	if err != nil {
		*cur++
		return nil
	}
	newNode := TreeNode{val, nil, nil}

	*cur++

	newNode.Left = this.deserializeRecur(strS, cur)
	newNode.Right = this.deserializeRecur(strS, cur)

	return &newNode
}

func main() {
	ser := Constructor()

	fmt.Println(ser.serialize(ser.deserialize("1 2 # # 3 4 # # 5 # #")))
}

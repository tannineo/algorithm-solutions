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
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	str := strconv.Itoa(root.Val)

	left := this.serialize(root.Left)
	right := this.serialize(root.Right)

	if left != "" {
		str += " " + left
	}
	if right != "" {
		str += " " + right
	}

	return str
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	root := new(TreeNode)

	strS := strings.Split(data, " ")
	for i, str := range strS {
		val, err := strconv.Atoi(str)
		if err != nil {
			return nil // invalid string (int), should not be here
		}

		// build BST
		if i == 0 {
			root.Val = val
		} else {
			cur := root
			for {
				if val > cur.Val {
					if cur.Right == nil {
						cur.Right = &TreeNode{val, nil, nil}
						break
					}
					cur = cur.Right
				} else if val < cur.Val {
					if cur.Left == nil {
						cur.Left = &TreeNode{val, nil, nil}
						break
					}
					cur = cur.Left
				} else {
					return nil // duplicate elements, should not be here
				}
			}
		}
	}

	return root
}

func main() {
	ser := Constructor()

	fmt.Println(ser.serialize(ser.deserialize("2 1 3 5 4 9 8 7 11")))
}

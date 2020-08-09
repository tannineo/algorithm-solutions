# 0270_closest-binary-search-tree-value

水一个二叉树遍历

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestValue(root *TreeNode, target float64) int {
    if root.Left == nil && root.Right == nil {
        return root.Val
    }

    if root.Left == nil && root.Right != nil {
        rightVal := closestValue(root.Right, target)
        if distance(rightVal, target) < distance(root.Val, target) {
            return rightVal
        }
        return root.Val
    }

    if root.Left != nil && root.Right == nil {
        leftVal := closestValue(root.Left, target)
        if distance(leftVal, target) < distance(root.Val, target) {
            return leftVal
        }
        return root.Val
    }


    rightVal := closestValue(root.Right, target)
    leftVal := closestValue(root.Left, target)
    if distance(rightVal, target) < distance(root.Val, target) && distance(rightVal, target) < distance(leftVal, target) {
            return rightVal
    }
    if distance(leftVal, target) < distance(root.Val, target) && distance(leftVal, target) < distance(rightVal, target) {
            return leftVal
    }

    return root.Val
}

func distance(aa int, b float64) float64 {
    a := float64(aa)
    if a - b < 0 {
        return b - a
    }
    return a - b
}
```

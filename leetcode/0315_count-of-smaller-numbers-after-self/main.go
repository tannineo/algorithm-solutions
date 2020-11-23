package main

type BSTNode struct {
	Val                              int
	SelfCount, LeftCount, RightCount int
	Left, Right                      *BSTNode
}

// BST
func countSmaller(nums []int) []int {
	ans := make([]int, len(nums))

	var bstRoot *BSTNode

	for i := len(nums) - 1; i >= 0; i-- {
		if bstRoot == nil {
			bstRoot = &BSTNode{
				Val:       nums[i],
				SelfCount: 1,
			}
			ans[i] = 0
			continue
		}
		ans[i] = countAndAddBST(nums[i], bstRoot)
	}

	return ans
}

func countAndAddBST(num int, root *BSTNode) int {
	cur := root
	sum := 0
	for {
		if cur.Val == num {
			cur.SelfCount++
			sum += cur.LeftCount
			break
		} else if cur.Val < num {
			sum += cur.LeftCount + cur.SelfCount
			cur.RightCount++
			if cur.Right == nil {
				cur.Right = &BSTNode{
					Val:       num,
					SelfCount: 1,
				}
				break
			} else {
				cur = cur.Right
			}
		} else { // cur.Val > num
			cur.LeftCount++
			if cur.Left == nil {
				cur.Left = &BSTNode{
					Val:       num,
					SelfCount: 1,
				}
				break
			} else {
				cur = cur.Left
			}
		}
	}

	return sum
}

// brute force
func countSmallerBF(nums []int) []int {
	ans := make([]int, len(nums))

	for i, num := range nums {
		ans[i] = 0
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < num {
				ans[i]++
			}
		}
	}

	return ans
}

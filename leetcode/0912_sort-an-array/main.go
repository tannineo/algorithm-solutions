package main

import "fmt"

func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	// quickSort(nums, 0, len(nums)-1)
	countSort(nums)

	return nums
}

// quickSort sorts the elements in nums in [l, r] (r included)
// time: O(nlogn) ~ O(n^2)
// space: O(logn) ~ O(n)
func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}

	i, j := l+1, r

	for i <= j {
		if nums[i] <= nums[l] {
			i++
		} else if nums[j] >= nums[l] {
			j--
		} else {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	nums[j], nums[l] = nums[l], nums[j]

	quickSort(nums, l, j-1)
	quickSort(nums, j+1, r)
}

// countSort
// time: O(n)
// space: O(max(nums) - min(nums))
func countSort(nums []int) {
	maxNum, minNum := maxOfNums(nums), minOfNums(nums)

	counts := make([]int, maxNum-minNum+1)
	for _, v := range nums {
		counts[v-minNum]++
	}

	cur := 0
	for i, v := range counts {
		for j := 0; j < v; j++ {
			nums[cur] = i + minNum
			cur++
		}
	}
}

// heapSort, we did not use the implementation from 'container'
// we implemented the maxHeap by ourselves here
// time: O(n*logn)
// space: O(n)
func heapSort(nums []int) {
	// init heap
	for i := len(nums) / 2; i >= 0; i-- {
		heapify(nums, i, len(nums)-1)
	}

	// get max to the end
	for i := len(nums) - 1; i >= 1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, 0, i-1)
	}
}

// heapify the nums [i, e] (e included)
func heapify(nums []int, i, e int) {
	for i <= e {
		l := 2*i + 1 // left child
		r := 2*i + 2 // right child
		j := i
		if l <= e && nums[l] > nums[j] {
			j = l
		}
		if r <= e && nums[r] > nums[j] {
			j = r
		}
		if j == i {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i = j
	}
}

// mergeSort sort the elements in nums in [l, r) (r not included)
// time: O(n*logn)
// space: O(logn + n)
func mergeSort(nums []int, l, r int) {
	if l+1 >= r {
		return
	}

	mid := l + (r-l)/2

	mergeSort(nums, l, mid)
	mergeSort(nums, mid, r)

	i, j := l, mid
	cur := 0
	temps := make([]int, r-l)

	for i < mid || j < r {
		if j == r || (i < mid && nums[i] < nums[j]) {
			temps[cur] = nums[i]
			i++
		} else {
			temps[cur] = nums[j]
			j++
		}
		cur++
	}

	for i, v := range temps {
		nums[l+i] = v
	}
}

// BST
// time: O(n*logn)
// space: O(n)
type TreeNode struct {
	Val, Count  int
	Left, Right *TreeNode
}

func buildBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	newBST := &TreeNode{nums[0], 1, nil, nil}

	for i := 1; i < len(nums); i++ {
		cur := newBST
		for {
			if cur.Val < nums[i] {
				if cur.Right == nil {
					cur.Right = &TreeNode{nums[i], 1, nil, nil}
					break
				} else {
					cur = cur.Right
				}
			} else if nums[i] < cur.Val {
				if cur.Left == nil {
					cur.Left = &TreeNode{nums[i], 1, nil, nil}
					break
				} else {
					cur = cur.Left
				}
			} else {
				cur.Count++
				break
			}
		}
	}

	return newBST
}

func toSortedArray(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	leftArr := toSortedArray(root.Left)

	midArr := []int{}
	for i := 0; i < root.Count; i++ {
		midArr = append(midArr, root.Val)
	}

	rightArr := toSortedArray(root.Right)

	return append(leftArr, append(midArr, rightArr...)...)
}

// ----------------------------------------------------------------
func maxOfNums(nums []int) int {
	ans := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > ans {
			ans = nums[i]
		}
	}

	return ans
}

func minOfNums(nums []int) int {
	ans := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < ans {
			ans = nums[i]
		}
	}

	return ans
}

func main() {
	fmt.Println(toSortedArray(buildBST([]int{4, 3, 2, 1, 5, 6, 3, 7, 8, 9, 4, 5, 7, 3, 5, 1, 2, 3})))
}

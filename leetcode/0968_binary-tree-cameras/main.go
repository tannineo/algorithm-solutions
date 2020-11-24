package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func min(arg ...int) int {
	if len(arg) == 0 {
		return 0
	}
	minVal := arg[0]
	for i := 1; i < len(arg); i++ {
		if minVal > arg[i] {
			minVal = arg[i]
		}
	}
	return minVal
}

func minCameraCover(root *TreeNode) int {
	// root != nil

	_, covered, hasCamera := minCameraCoverPossible(root)

	if covered == -1 { // when only one node
		return hasCamera
	}
	return min(covered, hasCamera)
}

func minCameraCoverPossible(root *TreeNode) (notCovered, covered, hasCamera int) {
	if root == nil { //
		return -1, -1, -1
	}
	if root.Left == nil && root.Right == nil {
		return 0, -1, 1
	}

	// xxxHasCamera will always >= 0 when xxx != nil
	// xxxNotCovered and xxxCovered might be invalid with -1
	leftNotCovered, leftCovered, leftHasCamera := minCameraCoverPossible(root.Left)
	rightNotCovered, rightCovered, rightHasCamera := minCameraCoverPossible(root.Right)

	// notCovered
	notCovered = -1
	if root.Left == nil {
		notCovered = rightCovered
	} else if root.Right == nil {
		notCovered = leftCovered
	} else { // both left and right != nil
		if leftCovered >= 0 && rightCovered >= 0 {
			notCovered = leftCovered + rightCovered
		} // else is not valid
	}

	// covered
	covered = -1
	if root.Left == nil {
		covered = rightHasCamera
	} else if root.Right == nil {
		covered = leftHasCamera
	} else { // both left and right != nil
		covered = leftHasCamera + rightHasCamera
		if leftCovered >= 0 {
			covered = min(covered, leftCovered+rightHasCamera)
		}
		if rightCovered >= 0 {
			covered = min(covered, leftHasCamera+rightCovered)
		}
	}

	// hasCamera
	hasCamera = -1
	if root.Left == nil {
		hasCamera = rightHasCamera
		if rightNotCovered >= 0 {
			hasCamera = min(hasCamera, rightNotCovered)
		}
		if rightCovered >= 0 {
			hasCamera = min(hasCamera, rightCovered)
		}
		hasCamera++
	} else if root.Right == nil {
		hasCamera = leftHasCamera
		if leftNotCovered >= 0 {
			hasCamera = min(hasCamera, leftNotCovered)
		}
		if leftCovered >= 0 {
			hasCamera = min(hasCamera, leftCovered)
		}
		hasCamera++
	} else { // both left and right != nil
		leftMin := leftHasCamera
		if leftNotCovered >= 0 {
			leftMin = min(leftMin, leftNotCovered)
		}
		if leftCovered >= 0 {
			leftMin = min(leftMin, leftCovered)
		}
		rightMin := rightHasCamera
		if rightNotCovered >= 0 {
			rightMin = min(rightMin, rightNotCovered)
		}
		if rightCovered >= 0 {
			rightMin = min(rightMin, rightCovered)
		}
		hasCamera = leftMin + rightMin + 1
	}

	return
}

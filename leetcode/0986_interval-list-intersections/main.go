package main

func intervalIntersection(A [][]int, B [][]int) [][]int {

	lenA := len(A)
	lenB := len(B)
	if lenA == 0 || lenB == 0 {
		return [][]int{}
	}

	result := [][]int{}

	curA := 0
	curB := 0
	inA := false
	inB := false
	curResult := -1

	for curA < lenA*2 && curB < lenB*2 {
		// 1. next is in A or B or BOTH ?
		if A[curA/2][curA%2] < B[curB/2][curB%2] {
			cursor := A[curA/2][curA%2]
			if inA && inB {
				result[curResult][1] = cursor
			} else if !inA && inB {
				result = append(result, []int{-1, -1})
				curResult++
				result[curResult][0] = cursor
			} else if inA && !inB {

			}
			inA = !inA
			curA++
		} else if A[curA/2][curA%2] > B[curB/2][curB%2] {
			cursor := B[curB/2][curB%2]
			if inA && inB {
				result[curResult][1] = cursor
			} else if inA && !inB {
				result = append(result, []int{-1, -1})
				curResult++
				result[curResult][0] = cursor
			}
			inB = !inB
			curB++
		} else { // ==
			cursor := A[curA/2][curA%2]
			if inA && inB {
				result[curResult][1] = cursor
			} else if !inA && !inB {
				result = append(result, []int{-1, -1})
				curResult++
				result[curResult][0] = cursor
			} else { // inA != inB
				result = append(result, []int{cursor, cursor})
				curResult++
			}
			inA = !inA
			inB = !inB
			curA++
			curB++
		}
	}

	return result
}

func main() {}

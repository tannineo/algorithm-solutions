package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubarraySumCircular(A []int) int {

	lenA := len(A)

	if lenA == 1 {
		return A[0]
	}

	dpLast := A[0]
	dpLastStart := 0
	dpAnswer := A[0]
	for i := 2; i <= lenA; i++ {
		if dpLast < 0 {
			dpLastStart = i - 1
			dpLast = 0 + A[i-1]
		} else {
			dpLast = dpLast + A[i-1]
		}
		if dpAnswer < dpLast {
			dpAnswer = dpLast
		}
	}

	var maxSumPreffix int
	sumPreffix := 0
	answer := dpAnswer
	for i := 0; i < lenA-1; i++ {
		for i+1 >= dpLastStart {
			// right shift
			dpLast -= A[dpLastStart]
			dpLastStart++
		}
		sumPreffix += A[i]
		if i == 0 {
			maxSumPreffix = sumPreffix
		} else {
			maxSumPreffix = max(maxSumPreffix, sumPreffix)
		}

		answer = max(answer, dpLast+maxSumPreffix)
	}

	return answer
}

func main() {
	println(maxSubarraySumCircular([]int{1, -2, 3, -2})) // 3

	println(maxSubarraySumCircular([]int{5, -3, 5})) // 10

	println(maxSubarraySumCircular([]int{3, -1, 2, -1})) // 4

	println(maxSubarraySumCircular([]int{3, -2, 2, -3})) // 3

	println(maxSubarraySumCircular([]int{-2, -3, -1})) // -1

	println(maxSubarraySumCircular([]int{-5, -2, 5, 6, -2, -7, 0, 2, 8})) // 14

	println(maxSubarraySumCircular([]int{1, -2, 3, -2})) // 3

	println(maxSubarraySumCircular([]int{5, -3, 5})) // 10
}

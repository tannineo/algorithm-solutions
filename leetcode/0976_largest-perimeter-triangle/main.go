package main

import "sort"

func largestPerimeter(A []int) int {

	lenA := len(A)

	// if FP, deep copy A to newA
	newA := make([]int, len(A))
	copy(newA, A)

	sort.Sort(sort.Reverse(sort.IntSlice(newA)))

	// let a <= b <= c, if a, b, c are the answer, they must satisfy a + b > c to form a triangle
	// if in the sorted array/slice [.... | a, b, c], we pick the known number `c`, and a + b <= c,
	// there will be no other numbers can satisfy the condition, so we look for the next value (here is `b`)

	for i := range newA {
		if i >= lenA-2 {
			break
		}
		if newA[i] < newA[i+1]+newA[i+2] {
			return newA[i] + newA[i+1] + newA[i+2]
		}
	}

	return 0
}

func main() {
	println(largestPerimeter([]int{3, 2, 8, 1}))

	println(largestPerimeter([]int{2, 1, 2}))

	println(largestPerimeter([]int{1, 2, 1}))

	println(largestPerimeter([]int{3, 2, 3, 4}))

	println(largestPerimeter([]int{3, 6, 2, 3}))

}

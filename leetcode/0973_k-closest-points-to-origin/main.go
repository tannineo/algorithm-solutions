package main

import "fmt"

func kClosest(points [][]int, K int) [][]int {
	lenP := len(points)

	// [0, lenP)
	partialSort(points, 0, lenP, K)

	return points[:K]
}

// MaxInt32 = 2147483647
func calcEu(points []int) int {
	return points[0]*points[0] + points[1]*points[1]
}

func partialSort(points [][]int, oi, oj, K int) {
	if K == oj-oi {
		return
	}
	// oi as pivot, [oi+1, oj)
	i := oi + 1
	j := oj - 1

	for {
		for i < oj-1 && calcEu(points[i]) <= calcEu(points[oi]) {
			i++
		}
		for j > oi && calcEu(points[j]) >= calcEu(points[oi]) {
			j--
		}
		if j <= i {
			break
		}
		points[i], points[j] = points[j], points[i]
	}

	points[oi], points[j] = points[j], points[oi]

	if j-oi < K {
		partialSort(points, j+1, oj, K-j+oi-1)
	} else {
		partialSort(points, oi, j, K)
	}
}

func main() {
	fmt.Printf("%v\n", kClosest([][]int{
		{-41, 72},
		{53, 83},
		{-95, -31},
		{-61, 68},
		{32, -56},
		{16, 88},
		{-81, -48},
		{-31, 56},
		{-57, -74},
		{24, -42},
		{-59, 72},
		{-86, 40},
		{34, -85},
		{-45, 22},
		{-35, -95},
	}, 9))
}

package main

func singleNumber(nums []int) (result int) {
	result = 0

	for _, num := range nums {
		result ^= num
	}

	return
}

func main() {
	println(singleNumber([]int{2, 2, 1})) // 1

	println(singleNumber([]int{4, 1, 2, 1, 2})) // 4
}

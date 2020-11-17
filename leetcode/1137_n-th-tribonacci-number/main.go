package main

import "fmt"

func tribonacci(n int) int {
	ans := []int{
		0,
		1,
		1,
		2,
		4,
		7,
		13,
		24,
		44,
		81,
		149,
		274,
		504,
		927,
		1705,
		3136,
		5768,
		10609,
		19513,
		35890,
		66012,
		121415,
		223317,
		410744,
		755476,
		1389537,
		2555757,
		4700770,
		8646064,
		15902591,
		29249425,
		53798080,
		98950096,
		181997601,
		334745777,
		615693474,
		1132436852,
		2082876103,
	}

	return ans[n]
}

func gen(n int) int {
	ans := []int{0, 1, 1}

	cur := 3

	for cur <= n {
		ans = append(ans, ans[cur-3]+ans[cur-2]+ans[cur-1])

		// print
		fmt.Println(ans[cur])

		cur++
	}

	return ans[n]
}

func main() {
	gen(37)
}

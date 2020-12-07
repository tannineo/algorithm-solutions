package main

import "fmt"

func consecutiveNumbersSum(num int) int {
	ans := 0

	for i := 1; i*(i-1)/2 <= num; i++ {
		rect := num - i*(i-1)/2
		if rect > 0 && rect%i == 0 { // 0 is not allowed in the Consecutive Numbers
			// fmt.Println("i=", i, "rect=", rect)
			ans++
		}
	}

	return ans
}

func main() {
	fmt.Println(consecutiveNumbersSum(5)) // 2

	fmt.Println(consecutiveNumbersSum(9)) // 3

	fmt.Println(consecutiveNumbersSum(15)) // 4
}

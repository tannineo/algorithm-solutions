package main

import "strconv"

func fizzBuzz(n int) []string {
	ans := make([]string, n)
	for i := range ans {
		num := i + 1
		if num%3 == 0 && num%5 == 0 {
			ans[i] = "FizzBuzz"
		} else if num%3 == 0 {
			ans[i] = "Fizz"
		} else if num%5 == 0 {
			ans[i] = "Buzz"
		} else {
			ans[i] = strconv.Itoa(num)
		}
	}

	return ans
}

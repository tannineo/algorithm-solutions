package main

import "fmt"

func checkValidString(s string) bool {

	if s == "" {
		return true
	}

	lenS := len(s)

	dp := make([][]bool, lenS+1)
	for i := range dp {
		dp[i] = make([]bool, lenS+1)
	}

	// init dp value
	for i := 0; i <= lenS; i++ {
		dp[i][i] = true
	}
	// calculate dp[i][j], 0 <= i < lenS, i <= j <= lenS
	// j = i + l
	// 0 <= l <= lenS
	for l := 1; l <= lenS; l++ {
		for i, _ := range s {
			if i+l > lenS { // end for lenth = l
				break
			}

			if s[i+l-1] == '*' && dp[i][i+l-1] {
				dp[i][i+l] = true
				continue
			}

			if s[i+l-1] != '(' {
				flag := false

				for k := i; k < i+l-1; k++ {
					if s[k] == ')' {
						continue
					}
					if dp[i][k] && dp[k+1][i+l-1] {
						flag = true
						break
					}
				}

				dp[i][i+l] = flag
			}

		}

	}

	return dp[0][lenS]
}

func checkValidStringGreedy(s string) bool {
	lower, higher := 0, 0

	for _, v := range s {
		if v == '(' {
			lower++
			higher++
		} else if v == ')' {
			lower--
			higher--
		} else { // == '*'
			lower--
			higher++
		}

		if higher < 0 {
			break
		}

		if lower < 0 {
			lower = 0
		}
	}

	return lower == 0
}

// 不快指数 根据温度湿度计算不快指数
// 温度为摄氏度
// 湿度为0-100的数字, 舍去百分号
func 不快指数(温度, 湿度 float64) float64 {
	return 0.81*温度 + 0.01*湿度*(0.99*温度-14.3) + 46.3
}

func main() {

	println(checkValidString("")) // true

	println(checkValidString("()*")) // true

	println(checkValidString("()")) // true

	println(checkValidString("(*)")) // true

	println(checkValidString("(*))")) // true

	println(checkValidString("*(()()")) // false

	println(checkValidString("(())((())()()(*)(*()(())())())()()((()())((()))(*")) // false

	println(checkValidStringGreedy("")) // true

	println(checkValidStringGreedy("()*")) // true

	println(checkValidStringGreedy("()")) // true

	println(checkValidStringGreedy("(*)")) // true

	println(checkValidStringGreedy("(*))")) // true

	println(checkValidStringGreedy("*(()()")) // false

	println(checkValidStringGreedy("(())((())()()(*)(*()(())())())()()((()())((()))(*")) // false

	fmt.Printf("某市: %v\n", 不快指数(35, 73))
	fmt.Printf("某郡: %v\n", 不快指数(32, 79))
	fmt.Printf("某道: %v\n", 不快指数(29, 91))
	fmt.Printf("又某市: %v\n", 不快指数(18, 52))
}

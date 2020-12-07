package main

import "fmt"

func numDecodings(s string) int {
	mapMemo := map[string]int{}
	return numDecodingsWithMemorization(s, mapMemo)
}

func numDecodingsWithMemorization(s string, mapMemo map[string]int) int {
	if v, ok := mapMemo[s]; ok {
		return v
	}

	if len(s) == 0 {
		return 1
	}
	if s[0] == '0' {
		return 0
	}
	if len(s) == 1 {
		return 1
	}

	if len(s) == 2 {
		if s[0] == '1' {
			if s[1] != '0' {
				return 2
			} else {
				return 1
			}
		} else if s[0] == '2' {
			if s[1] == '0' {
				return 1
			} else if s[1] <= '6' {
				return 2
			} else {
				return 1
			}
		} else { // 3X - 9X
			if s[1] == '0' {
				return 0
			}
			return 1
		}
	}

	// divide and conquer
	div := len(s) / 2
	ans := numDecodingsWithMemorization(s[:div], mapMemo) * numDecodingsWithMemorization(s[div:], mapMemo)
	if s[div-1] == '1' || (s[div-1] == '2' && s[div] <= '6') {
		ans += numDecodingsWithMemorization(s[:div-1], mapMemo) * numDecodingsWithMemorization(s[div+1:], mapMemo)
	}

	mapMemo[s] = ans // memorize

	return ans
}

func main() {
	fmt.Println(numDecodings("1111111111111111111111111"))
}

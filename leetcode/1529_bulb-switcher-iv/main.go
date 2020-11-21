package main

func minFlips(target string) int {
	flag := true
	curType := '1'
	ans := 0
	for _, v := range target {

		if flag {
			if v == '1' {
				flag = false
				curType = v
				ans++
			} else { // v == '0'
				continue
			}
		} else {
			if curType != v {
				ans++
				curType = v
			}
		}
	}

	return ans
}

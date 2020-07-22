package main

func rangeBitwiseAnd(m int, n int) int {
	if m == n {
		return m
	}

	bitsM := getBitCount(m)
	bitsN := getBitCount(n)

	if bitsM != bitsN {
		return 0
	}

	result := 0
	cur := 1 << (bitsM - 1) // bitsM won't be 0 here

	for cur > 0 {
		if m&cur == n&cur {
			result += (cur & m)
			cur >>= 1
		} else {
			break
		}
	}

	return result
}

func getBitCount(n int) int {
	result := 0
	for n > 0 {
		n >>= 1
		result += 1
	}
	return result
}

func main() {
	println(rangeBitwiseAnd(5, 7)) // 4
}

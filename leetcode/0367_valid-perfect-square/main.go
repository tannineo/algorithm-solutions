package main

func isPerfectSquareBS(num int) bool {
	if num == 1 || num == 4 || num == 9 {
		return true
	}
	if num < 9 {
		return false
	}

	left := 4
	right := num - 1
	// [left, guess), [guess, right)
	for left < right {
		guess := (left + right) / 2
		guessSqr := guess * guess
		if guessSqr == num {
			return true
		}

		if guessSqr > num {
			right = guess
		} else { // guessSqr < num
			left = guess + 1
		}
	}

	return false
}

// isPerfectSquare uses Newton's method
func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}

	x := num
	y := num
	for x >= y {
		x = y
		y = (x + num/x) / 2
		if x == y {
			return y*y == num
		}
	}

	return false
}

func main() {
	println(isPerfectSquare(16)) // true

	println(isPerfectSquare(14)) // true

}

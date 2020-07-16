package main

func isHappy(n int) bool {
	records := map[int]struct{}{}

	for n != 1 {
		if _, ok := records[n]; ok {
			return false
		} else {
			records[n] = struct{}{}
		}

		n = step(n)
		println(n)

	}

	return true
}

func step(num int) int {
	result := 0

	for num > 0 {
		digit := num % 10
		result += digit * digit
		num /= 10
	}

	return result
}

func main() {

	println(isHappy(19))
	println(isHappy(7))
}

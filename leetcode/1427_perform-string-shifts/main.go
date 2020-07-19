package main

func stringShift(s string, shifts [][]int) string {

	// 0. "zero" cases
	if len(shifts) == 0 {
		return s
	}

	// 1. sum up all the shift commands
	totalShift := 0
	for _, sh := range shifts {
		if sh[0] == 0 {
			totalShift -= sh[1]
		} else {
			totalShift += sh[1]
		}
	}

	// 2. perform the overall shift
	lenS := len(s)
	var splitPos int
	if totalShift < 0 {
		totalShift = -((-totalShift) % lenS)
		splitPos = -totalShift
	} else {
		totalShift = totalShift % lenS
		splitPos = lenS - totalShift
	}

	if totalShift == 0 {
		return s
	}

	runes := []rune(s)
	runes = append(runes[splitPos:], runes[:splitPos]...)

	return string(runes)
}

func main() {
	println(stringShift("abc", [][]int{
		[]int{0, 1},
		[]int{1, 2},
	}))
}

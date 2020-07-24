package main

func removeKdigits(num string, k int) string {
	runes := []rune(num)
	for k > 0 {
		cur := 0
		for cur+1 < len(runes) && runes[cur] <= runes[cur+1] {
			cur++
		}

		// delete num[cur]
		runes = append(runes[0:cur], runes[cur+1:]...)

		k--
	}

	// delete prefix zeros
	for len(runes) != 0 && runes[0] == '0' {
		runes = runes[1:]
	}

	if len(runes) == 0 {
		return "0"
	}

	newNum := string(runes)
	return newNum
}

func main() {

}

package main

func reverseWords(s string) string {
	lenS := len(s)
	if lenS == 0 {
		return s
	}

	// 1. remove redundant spaces
	lastIsWord := false
	nextToReplace := 0
	runes := []rune(s)
	for _, r := range runes {
		if r == ' ' && lastIsWord {
			lastIsWord = false
			runes[nextToReplace] = r
			nextToReplace++
		}
		if r != ' ' {
			lastIsWord = true
			runes[nextToReplace] = r
			nextToReplace++
		}
	}
	// rid of trailing space
	runes = runes[:nextToReplace]
	lenR := len(runes)
	if lenR == 0 {
		return ""
	}
	if runes[len(runes)-1] == ' ' {
		runes = runes[:len(runes)-1]
	}

	lenR = len(runes)
	if lenR == 0 {
		return ""
	}

	// 2. reverse whole string
	reverseRunes(runes, 0, lenR)

	// 3. reverse single words
	nextToReverse := 0
	for i, r := range runes {
		if i == lenR-1 {
			reverseRunes(runes, nextToReverse, lenR)
			break
		}
		if r == ' ' {
			reverseRunes(runes, nextToReverse, i)
			nextToReverse = i + 1
		}
	}

	return string(runes)
}

func reverseRunes(runes []rune, left, right int) {
	for i, j := left, right-1; i < j; {

		runes[i], runes[j] = runes[j], runes[i]

		i++
		j--
	}
}

func main() {
	println(reverseWords("  a good   example    s ") + "|")
}

package main

func getNextValidChar(str []rune, right int, skips int) (char rune, at int, skipsLeft int) {
	char = 0
	at = -1
	skipsLeft = skips

	for cursor := right - 1; cursor >= 0; cursor-- {
		if str[cursor] == '#' {
			skipsLeft++
		} else { // not '#'
			if skipsLeft == 0 {
				char = str[cursor]
				at = cursor
				return
			} else {
				skipsLeft--
			}
		}
	}

	return
}

func backspaceCompare(S string, T string) bool {

	sRunes := []rune(S)
	tRunes := []rune(T)

	curS := len(S)
	curT := len(T)

	skipsS := 0
	skipsT := 0

	for curS >= 0 && curT >= 0 {
		var curSChar, curTChar rune

		curSChar, curS, skipsS = getNextValidChar(sRunes, curS, skipsS)
		curTChar, curT, skipsT = getNextValidChar(tRunes, curT, skipsT)

		if curSChar != curTChar {
			return false
		}
	}

	return true
}

func main() {
	println(backspaceCompare("ab#c", "ad#c")) // true

	println(backspaceCompare("ab##", "c#d#")) // true

	println(backspaceCompare("a##c", "#a#c")) // true

	println(backspaceCompare("a#c", "b")) // false

}

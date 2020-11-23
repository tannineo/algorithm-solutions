package main

func isValid(s string) bool {
	// len(s) > 0

	runesS := []rune(s)

	stack := make([]rune, len(runesS))
	cur := -1

	for _, v := range runesS {
		if cur == -1 {
			if v == '(' || v == '[' || v == '{' {
				cur++
				stack[cur] = v
			} else {
				return false
			}
		} else {
			if v == ')' {
				if stack[cur] == '(' {
					cur--
				} else {
					return false
				}
			} else if v == ']' {
				if stack[cur] == '[' {
					cur--
				} else {
					return false
				}
			} else if v == '}' {
				if stack[cur] == '{' {
					cur--
				} else {
					return false
				}
			} else {
				cur++
				stack[cur] = v
			}
		}
	}

	return cur == -1
}

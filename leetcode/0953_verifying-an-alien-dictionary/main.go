package main

func isAlienSorted(words []string, order string) bool {
	if len(words) <= 1 {
		return true
	}

	toOrigin := [26]byte{}
	for i := range order {
		toOrigin[int(order[i]-'a')] = byte('a' + i)
	}

	// replace, back to normal
	for i := range words {
		byteS := []byte(words[i])
		for j := range byteS {
			byteS[j] = toOrigin[byteS[j]-'a']
		}
		words[i] = string(byteS)
	}

	for i := 0; i < len(words)-1; i++ {
		if words[i] > words[i+1] {
			return false
		}
	}

	return true
}

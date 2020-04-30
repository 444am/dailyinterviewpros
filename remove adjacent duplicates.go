func removeDuplicates(S string) string {
	byteS := []byte(S)
	i := 0
	removeAgain:
	length := len(byteS)
	for i < length {
		if i + 1 < length && byteS[i] == byteS[i+1] {
			byteS = append(byteS[:i], byteS[i+2:]...)
			if i > 0 {
				i--
			}
			goto removeAgain
		}
		i++
	}
	return string(byteS)
}

func removeDuplicates(s string, k int) string {
	byteS := []byte(s)
	i := 0
	removeAgain:
	length := len(byteS)
	for i < length {
		char := byteS[i]
		if i + k- 1 < length {
			for j := i+1; j <= i+k-1; j++ {
				if char != byteS[j] {
					goto nextChar
				}
			}
			byteS = append(byteS[:i], byteS[i+k:]...)
			if i > 0 {
				i = i - k + 1
                if i < 0 {
                    i = 0
                }
			}
			goto removeAgain
		}
		nextChar:
		i++
	}
	return string(byteS)
}
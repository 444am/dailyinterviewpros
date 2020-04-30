package main

func commonChars(A []string) []string {
	ret := []string{}
	length := len(A)
	compareList := make([]map[rune]int, length)
	for i := 0; i < length; i++ {
		tmpMap := make(map[rune]int)
		str := A[i]
		for _, char := range str {
			tmpMap[char]++
		}
		compareList[i] = tmpMap
	}
	base := compareList[0]
	for char, _ := range base {
		for base[char] > 0 {
			for i := 1; i < length; i++ {
				comparedMap := compareList[i]
				if num, ok := comparedMap[char]; ok && num > 0 {
					comparedMap[char]--
				} else {
					goto nextChar
				}
			}
			ret = append(ret, string(char))
			base[char]--
		}
	nextChar:
	}
	return ret
}

package main

/*
	key of KMP: figure out the meaning of PMT(Partial Match Table)

	the longest match of a word's prefix and suffix

	the value of each element in "next" table is PMT value of the letter one step in front of the element

	thus PMT[0] will always be 0 (no prefix & suffix), next[0] will always be -1 (for coding convenience) & next[1] will always be 0
*/

func getNext(p string) []int {
	j := -1
	i := 0
	next := make([]int, len(p))
	next[0] = j

	for i < len(p)-1 {
		if j == -1 || p[i] == p[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}

	return next
}

func KMP(s string, p string) bool {
	i := 0
	j := 0
	next := getNext(p)

	for i < len(s) && j < len(p) {
		if j == -1 || s[i] == p[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}

	return j == len(p)
}

func rotateString(A string, B string) bool {
	if len(A) == 0 && len(B) == 0 {
		return true
	}

	if len(A) != len(B) {
		return false
	}
	seamed := A + A
	return KMP(seamed, B)
}

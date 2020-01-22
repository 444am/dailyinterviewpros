package main

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	diff := 0
	indices := [2]int{}
	replication := make(map[byte]int)
	for i := 0; i < len(A); i++ {
		if _, ok := replication[A[i]]; ok {
			replication[A[i]]++
		} else {
			replication[A[i]] = 1
		}
		if A[i] != B[i] {
			if diff >= 2 {
				return false
			}
			indices[diff] = i
			diff++
		}
	}
	if diff == 0 {
		for _, v := range replication {
			if v >= 2 {
				return true
			}
		}
		return false
	}
	if A[indices[0]] == B[indices[1]] && A[indices[1]] == B[indices[0]] {
		return true
	}
	return false
}

package main

func heapify(i int, n int, words []*wordNfreq) {
	max := i
	left := i*2 + 1
	right := i*2 + 2

	if left < n {
		if words[left].freq > words[max].freq {
			max = left
		} else if words[left].freq == words[max].freq && words[left].word < words[max].word {
			max = left
		}
	}

	if right < n {
		if words[right].freq > words[max].freq {
			max = right
		} else if words[right].freq == words[max].freq && words[right].word < words[max].word {
			max = right
		}
	}

	if max != i {
		words[i], words[max] = words[max], words[i]
		heapify(max, n, words)
	}

	return
}

type wordNfreq struct {
	word string
	freq int
}

func topKFrequent(words []string, k int) []string {
	wordMap := make(map[string]int)

	for _, word := range words {
		wordMap[word]++
	}

	wnfList := []*wordNfreq{}

	for word, freq := range wordMap {
		wnfList = append(wnfList, &wordNfreq{word: word, freq: freq})
	}

	n := len(wnfList)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(i, n, wnfList)
	}

	res := []string{}
	for i := 0; i < k; i++ {
		res = append(res, wnfList[0].word)
		wnfList[0], wnfList[n-1-i] = wnfList[n-1-i], wnfList[0]
		heapify(0, n-1-i, wnfList)
	}

	return res
}

package main

func longestConsecutive(nums []int) int {
	set := make(map[int]struct{})
	res := 0

	for _, num := range nums {
		if _, ok := set[num]; !ok {
			set[num] = struct{}{}
		}
	}

	for num := range set {
		if _, ok := set[num-1]; !ok {
			tempLength := 1
			currentNum := num + 1
			for {
				if _, ok := set[currentNum]; ok {
					currentNum += 1
					tempLength += 1
				} else {
					break
				}
			}
			if tempLength > res {
				res = tempLength
			}
		}
	}

	return res
}

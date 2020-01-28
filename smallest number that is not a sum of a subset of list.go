package main

func findSmallest(nums []int) int {
	res := 1
	for _, num := range nums {
		if num <= res {
			res += num
		} else {
			break
		}
	}
	return res
}

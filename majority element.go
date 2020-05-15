package main

func majorityElement(nums []int) int {
	hashMap := make(map[int]int)
	threshold := len(nums) / 2
	ret := 0
	for _, num := range nums {
		hashMap[num]++
		if hashMap[num] > threshold {
			ret = num
			break
		}
	}
	return ret
}

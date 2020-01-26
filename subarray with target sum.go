package main

/*
	keep recording cumulative sum of current index & calculate sum - target int then record this value as the key of a hashmap and current index as its value
	if (sum - target int) ever appears in the hashmap, that means sum of subarray from relative index in the hashmap to current index equals target int
*/

func subarraySum(nums []int, k int) [][]int {
	hashMap := make(map[int][]int, len(nums))
	hashMap[0] = []int{-1}
	sum := 0
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if indices, ok := hashMap[sum-k]; ok {
			for _, index := range indices {
				res = append(res, nums[index+1:i+1])
			}
		}
		if _, ok := hashMap[sum]; ok {
			hashMap[sum] = append(hashMap[sum], i)
		} else {
			hashMap[sum] = []int{i}
		}
	}
	return res
}

package main

func sortColors(nums []int) {
	base := 1
	mid := 0
	left := 0
	right := len(nums) - 1
	for mid <= right {
		switch {
		case nums[mid] > base:
			nums[mid], nums[right] = nums[right], nums[mid]
			right--
		case nums[mid] < base:
			nums[mid], nums[left] = nums[left], nums[mid]
			left++
			mid++
		default:
			mid++
		}
	}
}

package main

func searchRange(nums []int, target int) []int {
    return []int{searchLeft(nums, target), searchRight(nums, target)}
}

func searchLeft(nums []int, target int) int {
	left := 0
	right := len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		switch {
		case nums[mid] < target:
			left = mid + 1
		case nums[mid] > target:
			right = mid - 1
		case nums[mid] == target:
			if mid == 0 || nums[mid-1] != target {
				return mid
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func searchRight(nums []int, target int) int {
	left := 0
	right := len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		switch {
		case nums[mid] < target:
			left = mid + 1
		case nums[mid] > target:
			right = mid - 1
		case nums[mid] == target:
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}
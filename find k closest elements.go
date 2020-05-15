package main

func findClosestElements(arr []int, k int, x int) []int {
	pivot := binarySearch(arr, x)
	k--
	leftBound := pivot
	rightBound := pivot
	for k > 0 {
		if leftBound-1 >= 0 {
			if rightBound+1 < len(arr) && abs(x, arr[rightBound+1]) < abs(x, arr[leftBound-1]) {
				rightBound++
			} else {
				leftBound--
			}
		} else {
			rightBound++
		}
		k--
	}
	return arr[leftBound : rightBound+1]
}

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)/2
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] > target:
			right = mid - 1
		case nums[mid] < target:
			left = mid + 1
		}
	}
	if left > 0 && abs(target, nums[left-1]) <= abs(target, nums[left]) {
		return left - 1
	}
	return left
}

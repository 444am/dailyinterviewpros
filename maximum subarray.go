package main

import "math"

func maxSubArray(nums []int) int {
	leftBound := 0
	rightBound := 0
	sum := math.MinInt32
	for leftBound < len(nums) {
		temp := 0
		broken := false
		for rightBound < len(nums) {
			temp += nums[rightBound]
			if temp > sum {
				sum = temp
			}
			if temp < 0 {
				leftBound = rightBound + 1
				rightBound = leftBound
				broken = true
				break
			}
			rightBound++
		}
		if !broken {
			leftBound++
		}
	}
	return sum
}

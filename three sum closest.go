package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	length := len(nums)
	diff := math.MaxInt32
	ret := 0
	for i := 0; i < len(nums)-2; i++ {
		for j, k := i+1, length-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			switch {
			case sum == target:
				return target
			case sum > target:
				k--
			default:
				j++
			}
			if abs(sum, target) < diff {
				diff = abs(sum, target)
				ret = sum
			}
		}
	}
	return ret
}

func abs(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

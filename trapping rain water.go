package main

import "fmt"

func trap(height []int) int {
	left := 0
	right := len(height) - 1
	leftMax := 0
	rightMax := 0
	res := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				res += leftMax - height[left]
			}
			left += 1
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				res += rightMax - height[right]
			}
			right -= 1
		}
	}

	return res
}

func main() {
	sample := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(sample))
}

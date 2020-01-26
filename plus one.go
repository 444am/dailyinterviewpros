package main

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		temp := digits[i] + carry
		if temp == 10 {
			temp = 0
			carry = 1
		} else {
			carry = 0
		}
		digits[i] = temp
	}
	if carry == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

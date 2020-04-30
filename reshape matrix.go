package main

func matrixReshape(nums [][]int, r int, c int) [][]int {
	origR := len(nums)
	origC := len(nums[0])

	if origC*origR != c*r {
		return nums
	}

	ret := [][]int{}
	p := 0
	for rPointer := 0; rPointer < r; rPointer++ {
		tempRow := []int{}
		for cPointer := 0; cPointer < c; cPointer++ {
			i := p / origC
			j := p % origC
			tempRow = append(tempRow, nums[i][j])
			p++
		}
		ret = append(ret, tempRow)
	}

	return ret
}

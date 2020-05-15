package main

func combinationSum(candidates []int, target int) [][]int {
	ret := [][]int{}
	helper(candidates, target, 0, []int{}, &ret)
	return ret
}

func helper(candidates []int, target int, startIndex int, temp []int, ret *[][]int) {
	if target == 0 {
		*ret = append(*ret, temp)
		return
	}
	if target < 0 {
		return
	}
	for i := startIndex; i < len(candidates); i++ {
		new := make([]int, len(temp)+1)
		copy(new, append(temp, candidates[i]))
		helper(candidates, target-candidates[i], i, new, ret)
	}
}

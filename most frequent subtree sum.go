package main

func findFrequentTreeSum(root *TreeNode) []int {
	ret := []int{}
	if root == nil {
		return ret
	}
	hashMap := make(map[int]int)
	maxF := 0
	dfsSum(root, hashMap, &maxF, &ret)
	return ret
}

func dfsSum(root *TreeNode, hashMap map[int]int, maxF *int, ret *[]int) int {
	if root == nil {
		return 0
	}
	left := dfsSum(root.Left, hashMap, maxF, ret)
	right := dfsSum(root.Right, hashMap, maxF, ret)
	sum := root.Val + left + right
	hashMap[sum]++
	if hashMap[sum] > *maxF {
		*maxF = hashMap[sum]
		*ret = []int{sum}
	} else if hashMap[sum] == *maxF {
		*ret = append(*ret, sum)
	}
	return sum
}

package main

func sumNumbers(root *TreeNode) int {
	ret := []int{}
	traversal(root, 0, &ret)
	sum := 0
	for _, v := range ret {
		sum += v
	}
	return sum
}

func traversal(root *TreeNode, sum int, ret *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*ret = append(*ret, root.Val+sum*10)
		return
	}
	traversal(root.Left, root.Val+sum*10, ret)
	traversal(root.Right, root.Val+sum*10, ret)
}

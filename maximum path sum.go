package main

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ret := 0
	sumHelper(root, &ret)
	return ret
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func sumHelper(root *TreeNode, ret *int) int {
	leftSum := 0
	rightSum := 0
	if root.Left != nil {
		leftSum = max(sumHelper(root.Left, ret), 0)
	}
	if root.Right != nil {
		rightSum = max(sumHelper(root.Right, ret), 0)
	}
	*ret = max(root.Val+leftSum+rightSum, *ret)
	return root.Val + max(leftSum, rightSum)
}

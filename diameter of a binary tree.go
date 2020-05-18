package main

func diameterOfBinaryTree(root *TreeNode) int {
	ret := 0
	depth(root, &ret)
	return ret
}

func depth(root *TreeNode, ret *int) int {
	if root == nil {
		return 0
	}
	left := depth(root.Left, ret)
	right := depth(root.Right, ret)
	if left+right > *ret {
		*ret = left + right
	}
	if left >= right {
		return left + 1
	}
	return right + 1
}

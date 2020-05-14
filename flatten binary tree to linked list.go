package main

func flatten(root *TreeNode) {
	dfs(root)
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	temp := root.Right
	root.Right = root.Left
	if root.Left != nil {
		dfs(root.Left)
	}
	root.Left = nil
	for root.Right != nil {
		root = root.Right
	}
	root.Right = temp
	dfs(root.Right)
}

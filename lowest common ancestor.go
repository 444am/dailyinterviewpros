package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	leftNode := lowestCommonAncestor(root.Left, p, q)
	rightNode := lowestCommonAncestor(root.Right, p, q)

	if leftNode == nil && rightNode == nil {
		return nil
	}

	if leftNode != nil && rightNode != nil {
		return root
	}

	if leftNode != nil {
		return leftNode
	}

	return rightNode
}

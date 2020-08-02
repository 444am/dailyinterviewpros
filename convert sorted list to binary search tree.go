package main

func sortedListToBST(head *ListNode) *TreeNode {
	node := head
	length := lengthofLinkedList(node)
	ret, _ := inorder(0, length-1, head)
	return ret
}

func lengthofLinkedList(head *ListNode) int {
	ret := 0
	for head != nil {
		ret++
		head = head.Next
	}
	return ret
}

func inorder(start int, end int, head *ListNode) (*TreeNode, *ListNode) {
	if start > end {
		return nil, nil
	}
	mid := (start + end) / 2
	left, node := inorder(start, mid-1, head)
	root := &TreeNode{Left: left, Val: node.Val}
	node = node.Next
	right, finalNode := inorder(mid+1, end, node)
	root.Right = right
	return root, finalNode
}

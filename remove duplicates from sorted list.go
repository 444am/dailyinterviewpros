package main

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	node := dummy
	for node != nil && node.Next != nil {
		val := node.Next.Val
		follow := node.Next.Next
		duplicate := false
		for follow != nil && follow.Val == val {
			duplicate = true
			follow = follow.Next
		}
		if duplicate {
			node.Next = follow
		} else {
			node = node.Next
		}
	}
	return dummy.Next
}

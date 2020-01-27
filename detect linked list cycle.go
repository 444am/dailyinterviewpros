package main

func hasCycle(head *ListNode) bool {
	nodeMap := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := nodeMap[head]; ok {
			return true
		} else {
			nodeMap[head] = struct{}{}
			head = head.Next
		}
	}
	return false
}

package main

func isPalindrome(head *ListNode) bool {
	fasterHead := head
	odd := false
	var dummy *ListNode
	dummy = nil

	for fasterHead != nil {
		fasterHead = fasterHead.Next
		if fasterHead != nil {
			fasterHead = fasterHead.Next
		} else {
			odd = true
			break
		}
		temp := head.Next
		head.Next = dummy
		dummy = head
		head = temp

	}

	if odd {
		head = head.Next
	}

	for dummy != nil {
		if dummy.Val != head.Val {
			return false
		} else {
			dummy = dummy.Next
			head = head.Next
		}
	}
	return true

}

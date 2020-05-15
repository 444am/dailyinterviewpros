package main

func validPalindrome(s string) bool {
	head := 0
	tail := len(s) - 1
	preservedHead := 0
	preservedTail := 0
	if decide(s, &head, &tail) {
		return true
	}
	preservedHead = head
	preservedTail = tail - 1

	head++
	if head <= tail {
		if decide(s, &head, &tail) {
			return true
		}
	}

	head = preservedHead
	tail = preservedTail
	if head <= tail {
		if decide(s, &head, &tail) {
			return true
		}
	}
	return false
}

func decide(s string, head *int, tail *int) bool {
	for *head <= *tail {
		if s[*head] != s[*tail] {
			return false
		}
		*head++
		*tail--
	}
	return true
}

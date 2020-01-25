package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	q []*TreeNode
}

func (q *Queue) Len() int {
	return len(q.q)
}

func (q *Queue) Queue(elem *TreeNode) {
	q.q = append(q.q, elem)
}

func (q *Queue) Dequeue() *TreeNode {
	ret := q.q[0]
	q.q = q.q[1:]
	return ret
}

type ListNode struct {
	Val  int
	Next *ListNode
}

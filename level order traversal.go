package main

import (
	"errors"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	queue []*TreeNode
}

func (q *Queue) Len() int {
	return len(q.queue)
}

func (q *Queue) Queue(e *TreeNode) {
	q.queue = append(q.queue, e)
}

func (q *Queue) Dequeue() (*TreeNode, error) {
	if len(q.queue) == 0 {
		return nil, errors.New("queue empty")
	}

	res := q.queue[0]
	q.queue = q.queue[1:len(q.queue)]
	return res, nil
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	queue := new(Queue)
	queue.Queue(root)
	for queue.Len() > 0 {
		temp := []int{}
		loopTimes := queue.Len()
		for i := 0; i < loopTimes; i++ {
			treenode, _ := queue.Dequeue()
			temp = append(temp, treenode.Val)
			if treenode.Left != nil {
				queue.Queue(treenode.Left)
			}
			if treenode.Right != nil {
				queue.Queue(treenode.Right)
			}
		}
		res = append(res, temp)
	}
	return res
}

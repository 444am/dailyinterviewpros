package main

import (
	"math"
)

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	q := new(Queue)
	q.Queue(root)
	max := math.MinInt64
	maxLevel := 0
	curLevel := 1
	for q.Len() > 0 {
		tempSum := 0
		execTimes := q.Len()

		for i := 0; i < execTimes; i++ {
			node := q.Dequeue()
			if node.Left != nil {
				q.Queue(node.Left)
			}
			if node.Right != nil {
				q.Queue(node.Right)
			}
			tempSum += node.Val
		}

		if tempSum > max {
			max = tempSum
			maxLevel = curLevel
		}

		curLevel += 1
	}

	return maxLevel
}

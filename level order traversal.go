package main

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
			treenode := queue.Dequeue()
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

package main

import (
	"strconv"
)

// serialize a tree: make a tree to a string, use # to occupy nil
// compare serialized tree

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	ret := []*TreeNode{}
	if root == nil {
		return ret
	}
	countMap := make(map[string]int)
	dfsSerialize(root, countMap, &ret)
	return ret
}

func dfsSerialize(root *TreeNode, countMap map[string]int, ret *[]*TreeNode) string {
	if root == nil {
		return "#"
	}
	serial := strconv.Itoa(root.Val) + " " + dfsSerialize(root.Left, countMap, ret) + " " + dfsSerialize(root.Right, countMap, ret)
	if _, ok := countMap[serial]; ok {
		countMap[serial]++
		if countMap[serial] == 2 {
			*ret = append(*ret, root)
		}
	} else {
		countMap[serial]++
	}
	return serial
}

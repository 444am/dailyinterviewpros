package main

type iNr struct {
	i int
	r int
}

func kClosest(points [][]int, K int) [][]int {
	ret := [][]int{}
	if len(points) == 0 {
		return ret
	}
	sortedList := []*iNr{}
	n := len(points)
	for i := 0; i < n; i++ {
		temp := &iNr{i: i, r: points[i][0]*points[i][0] + points[i][1]*points[i][1]}
		sortedList = append(sortedList, temp)
	}
	for i := n/2 - 1; i >= 0; i-- {
		heapifyiNr(i, n, sortedList)
	}
	for i := 0; i < K; i++ {
		ret = append(ret, points[sortedList[0].i])
		sortedList[0], sortedList[n-1-i] = sortedList[n-1-i], sortedList[0]
		heapifyiNr(0, n-1-i, sortedList)
	}
	return ret
}

func heapifyiNr(i int, n int, list []*iNr) {
	min := i
	left := i*2 + 1
	right := i*2 + 2
	if left < n && list[left].r < list[min].r {
		min = left
	}
	if right < n && list[right].r < list[min].r {
		min = right
	}
	if i != min {
		list[i], list[min] = list[min], list[i]
		heapifyiNr(min, n, list)
	}
	return
}

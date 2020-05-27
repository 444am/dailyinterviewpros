package main

func merge(intervals [][]int) [][]int {
	ret := [][]int{}
	quicksortInts(intervals, 0, len(intervals)-1)
	for _, interval := range intervals {
		if len(ret) == 0 || ret[len(ret)-1][1] < interval[0] {
			ret = append(ret, interval)
		} else {
			if ret[len(ret)-1][1] < interval[1] {
				ret[len(ret)-1][1] = interval[1]
			}
		}
	}
	return ret
}

func quicksortInts(intervals [][]int, left int, right int) {
	if left >= right {
		return
	}
	split := partitionInts(intervals, left, right)
	quicksortInts(intervals, left, split-1)
	quicksortInts(intervals, split+1, right)
}

func partitionInts(intervals [][]int, left int, right int) int {
	pivot := intervals[right][0]
	i := left - 1
	for j := left; j < right; j++ {
		if intervals[j][0] < pivot {
			i++
			intervals[i], intervals[j] = intervals[j], intervals[i]
		}
	}
	intervals[i+1], intervals[right] = intervals[right], intervals[i+1]
	return i + 1
}

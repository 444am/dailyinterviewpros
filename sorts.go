package main

func bubbleSort(list []int) {
	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-1-i; j++ {
			if list[j+1] < list[j] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}

func quicksort(list []int, left int, right int) {
	if left >= right {
		return
	}
	split := partition(list, left, right)
	quicksort(list, left, split-1)
	quicksort(list, split+1, right)
}

func partition(list []int, left int, right int) int {
	pivot := list[right]
	i := left - 1
	for j := left; j < right; j++ {
		if list[j] < pivot {
			i++
			list[i], list[j] = list[j], list[i]
		}
	}
	list[i+1], list[right] = list[right], list[i+1]
	return i + 1
}

func mergesort(list []int) {
	if len(list) <= 1 {
		return
	}
	mid := len(list) / 2
	left := make([]int, len(list[:mid]))
	right := make([]int, len(list[mid:]))
	copy(left, list[:mid])
	copy(right, list[mid:])
	mergesort(left)
	mergesort(right)
	i := 0
	j := 0
	k := 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			list[k] = left[i]
			i++
		} else {
			list[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		list[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		list[k] = right[j]
		j++
		k++
	}
}

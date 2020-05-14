package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row := len(matrix)
	cow := len(matrix[0])
	rowIndex := binarySearchRow(matrix, target, row, cow)
	if rowIndex == -1 {
		return false
	}
	return binarySearchCow(matrix[rowIndex], target, cow)
}

func binarySearchRow(matrix [][]int, target int, row int, cow int) int {
	left := 0
	right := row - 1
	for left <= right {
		mid := (left + right) / 2
		switch {
		case matrix[mid][0] == target:
			return mid
		case matrix[mid][0] > target:
			right = mid - 1
		case matrix[mid][0] < target:
			if matrix[mid][cow-1] >= target {
				return mid
			}
			left = mid + 1
		}
	}
	return -1
}

func binarySearchCow(row []int, target int, cow int) bool {
	left := 0
	right := cow - 1
	for left <= right {
		mid := (left + right) / 2
		switch {
		case row[mid] == target:
			return true
		case row[mid] > target:
			right = mid - 1
		case row[mid] < target:
			left = mid + 1
		}
	}
	return false
}

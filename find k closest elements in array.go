func findClosestElements(arr []int, k int, x int) []int {
	if x <= arr[0] {
		return arr[:k]
	}

	n := len(arr)
	if x >= arr[n-1] {
		return arr[n-k:]
	}
	pivot := binarySearchPivot(arr, x)
	left := max(0, pivot-k+1)
	right := min(n-1, pivot+k-1)
	for right-left > k-1 {
		if left < 0 || (x-arr[left]) <= (arr[right]-x) {
			right--
		} else {
			if right >= n || (x-arr[left]) > (arr[right]-x) {
				left++
			}
		}
	}
	return arr[left : right+1]
}

func binarySearchPivot(arr []int, x int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := left + (right-left)/2
		switch {
		case arr[mid] == x:
			return mid
		case arr[mid] > x:
			right = mid - 1
		default:
			left = mid + 1
		}
	}
	if abs(arr[left], x) < abs(arr[right], x) {
		return left
	}
	return right
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
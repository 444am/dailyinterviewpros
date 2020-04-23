package main

func transpose(A [][]int) [][]int {
	row := len(A)
	column := len(A[0])

	ret := [][]int{}
	for i := 0; i < column; i++ {
		newRow := []int{}
		for j := 0; j < row; j++ {
			newRow = append(newRow, A[j][i])
		}
		ret = append(ret, newRow)
	}

	return ret
}

package main

func generate(numRows int) [][]int {
	ret := [][]int{}
	for i := 1; i <= numRows; i++ {
		row := make([]int, i)
		row[0] = 1
		row[i-1] = 1
		for n := 1; n < i-1; n++ {
			row[n] = ret[i-2][n-1] + ret[i-2][n]
		}
		ret = append(ret, row)
	}
	return ret
}

func getRow(rowIndex int) []int {
	previousRow := []int{}
	for i := 0; i <= rowIndex+1; i++ { // index starts from 0, thus upperbound equals rowIndex + 1
		row := make([]int, i)
		row[0] = 1
		row[i-1] = 1
		for n := 1; n < i-1; n++ {
			row[n] = previousRow[n-1] + previousRow[n]
		}
		previousRow = row
	}
	return previousRow
}

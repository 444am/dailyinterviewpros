package main

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	is := intSlice(nums)
	sort.Sort(is)
	var buf bytes.Buffer
	for _, num := range is {
		buf.WriteString(strconv.Itoa(num))
	}
	temp := buf.String()

	for temp[0] == '0' && len(temp) > 1 {
		temp = temp[1:]
	}

	return temp
}

type intSlice []int

func (is intSlice) Len() int {
	return len(is)
}

func (is intSlice) Less(i, j int) bool {
	numI, _ := strconv.Atoi(strconv.Itoa(is[i]) + strconv.Itoa(is[j]))
	numJ, _ := strconv.Atoi(strconv.Itoa(is[j]) + strconv.Itoa(is[i]))
	//reverse sort
	return numI > numJ
}

func (is intSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func main() {
	a := []int{3, 30, 34, 5, 9}
	fmt.Println(largestNumber(a))
}

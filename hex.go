package main

func toHex(n int) string {
	ret := ``
	left := n
	decimalToHexMap := make(map[int]string)
	decimalToHexMap[0] = `0`
	decimalToHexMap[1] = `1`
	decimalToHexMap[2] = `2`
	decimalToHexMap[3] = `3`
	decimalToHexMap[4] = `4`
	decimalToHexMap[5] = `5`
	decimalToHexMap[6] = `6`
	decimalToHexMap[7] = `7`
	decimalToHexMap[8] = `8`
	decimalToHexMap[9] = `9`
	decimalToHexMap[10] = `A`
	decimalToHexMap[11] = `B`
	decimalToHexMap[12] = `C`
	decimalToHexMap[13] = `D`
	decimalToHexMap[14] = `E`
	decimalToHexMap[15] = `F`
	for left > 16 {
		mod := left % 16
		left = left / 16
		ret = decimalToHexMap[mod] + ret
	}
	ret = decimalToHexMap[left] + ret
	return ret
}

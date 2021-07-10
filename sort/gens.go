package sort

import "math/rand"

//GenArr is a func that generate array with length of "i" and the numbers are in maximum "max"
func GenArr(i int, max int) []int {
	arr := make([]int, i)
	for k := 0; k < i; k++ {
		arr[k] = rand.Intn(max)
	}
	return arr
}

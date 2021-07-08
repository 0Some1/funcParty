package main

import "math/rand"

func genArr(i int, max int) []int {
	arr := make([]int, i)
	for k := 0; k < i; k++ {
		arr[k] = rand.Intn(max)
	}
	return arr
}

package arrays

import "sync"

//MergeSort sort an array in a merge style
func MergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
			return arr
		}
		return arr
	}

	a := MergeSort(arr[:len(arr)/2])
	b := MergeSort(arr[len(arr)/2:])

	i := 0
	j := 0
	index := 0
	arr2 := make([]int, len(arr))
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			arr2[index] = a[i]
			i++
			index++
		} else {
			arr2[index] = b[j]
			j++
			index++
		}
	}

	if i == len(a) {
		for g := j; g < len(b); g++ {
			arr2[index] = b[g]
			index++
		}
	}
	if j == len(b) {
		for g := i; g < len(a); g++ {
			arr2[index] = a[g]
			index++
		}
	}

	return arr2
}

/*MergeSortC is the most important part ,
the function mergesortC take an array and number of cpu
and it will make go routine as many as it is efficient (base on the cpu cores that you have)
*/
func MergeSortC(arr []int, cpu int) []int {
	if len(arr) == 1 {
		return arr
	}
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
			return arr
		}
		return arr
	}

	var a []int
	var b []int
	wg := sync.WaitGroup{}
	if cpu >= 2 {
		wg.Add(2)
		cpu -= 2
		go func() {
			a = MergeSortC(arr[:len(arr)/2], cpu)
			wg.Done()
		}()
		go func() {
			b = MergeSortC(arr[len(arr)/2:], cpu)
			wg.Done()
		}()
		wg.Wait()
	} else {
		a = MergeSortC(arr[:len(arr)/2], cpu)

		b = MergeSortC(arr[len(arr)/2:], cpu)
	}

	i := 0
	j := 0
	index := 0
	arr2 := make([]int, len(arr))
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			arr2[index] = a[i]
			i++
			index++
		} else {
			arr2[index] = b[j]
			j++
			index++
		}
	}

	if i == len(a) {
		for g := j; g < len(b); g++ {
			arr2[index] = b[g]
			index++
		}
	}
	if j == len(b) {
		for g := i; g < len(a); g++ {
			arr2[index] = a[g]
			index++
		}
	}

	return arr2
}

/*
	this func failed and take much much more time than a simple sort algo
	because it does not consider the cpu cores
	and it will produce go routines as many as possible (require)
*/
//func mergeSortC(arr []int, c chan []int) {
//
//	if len(arr) == 1 {
//		c <- arr
//		return
//	} else if len(arr) == 2 {
//		if arr[0] > arr[1] {
//			arr[0], arr[1] = arr[1], arr[0]
//			c <- arr
//			return
//		} else {
//			c <- arr
//			return
//		}
//	}
//	var a []int
//	var b []int
//	c1 := make(chan []int)
//	c2 := make(chan []int)
//	go func() {
//		mergeSortC(arr[:len(arr)/2], c1)
//
//	}()
//
//	go func() {
//		mergeSortC(arr[len(arr)/2:], c2)
//
//	}()
//
//	a = <-c1
//	b = <-c2
//
//	i := 0
//	j := 0
//	index := 0
//	arr2 := make([]int, len(arr))
//	for i < len(a) && j < len(b) {
//		if a[i] <= b[j] {
//			arr2[index] = a[i]
//			i++
//			index++
//		} else {
//			arr2[index] = b[j]
//			j++
//			index++
//		}
//	}
//
//	if i == len(a) {
//		for g := j; g < len(b); g++ {
//			arr2[index] = b[g]
//			index++
//		}
//	}
//	if j == len(b) {
//		for g := i; g < len(a); g++ {
//			arr2[index] = a[g]
//			index++
//		}
//	}
//
//	c <- arr2
//	return
//}

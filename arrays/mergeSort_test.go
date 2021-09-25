package arrays

import (
	"runtime"
	"testing"
)

//BenchmarkMergeSortC is measuring the time of execution
func BenchmarkMergeSortC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := 1000000
		MergeSortC(GenArr(x, x), runtime.NumCPU())
	}
}

//BenchmarkMergeSort is measuring the time of execution
func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := 1000000
		MergeSort(GenArr(x, x))
	}
}

/*var inputLen int
fmt.Println("Enter an integer for array length: ")
_, err := fmt.Scan(&inputLen)
if err != nil {
	fmt.Println("error: ", err)
}

arr := GenArr(inputLen, inputLen)

arr2 := make([]int, len(arr))
arr3 := make([]int, len(arr))
arr4 := make([]int, len(arr))

for i := 0; i < len(arr); i++ {
	arr2[i] = arr[i]
	arr3[i] = arr[i]
	arr4[i] = arr[i]
}

fmt.Println(arr, "\n length of array: \t", len(arr), "\n")
fmt.Println("======================= normal Mergesort =================")

start1 := time.Now()
arr = MergeSort(arr)
end1 := time.Since(start1)
//fmt.Println(arr)
fmt.Println("\n time to sort an array in normal Mergesort:\t", end1, "\n")

fmt.Println("==================== concurrent Mergesort =================")

start3 := time.Now()
arr3 = MergeSortC(arr3, runtime.NumCPU())
end3 := time.Since(start3)
//fmt.Println(arr3)
fmt.Println("\n time to sort an array in concurrent mode:\t", end3, "\n")

fmt.Println("==================== golang sort  =================")

start4 := time.Now()
sort.Ints(arr4)
end4 := time.Since(start4)
//fmt.Println(arr4)
fmt.Println("\n time to sort an array with golang sort builtin function :\t", end4, "\n")
*/

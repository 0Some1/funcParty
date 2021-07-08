package main
import (
"fmt"
"runtime"
"sort"
"time"
)

func main() {
//BenchMark of three sort algorithms with measuring the time of execution
testMergeSort()
}

func testMergeSort()  {
arr := genArr(10000000,10000000)

arr2 := make([]int, len(arr))
arr3 := make([]int, len(arr))
arr4 := make([]int, len(arr))

for i := 0; i < len(arr); i++ {
arr2[i] = arr[i]
arr3[i] = arr[i]
arr4[i] = arr[i]
}

fmt.Println(arr, "\n length of array: \t", len(arr),"\n")
fmt.Println("======================= normal Mergesort =================")

start1 := time.Now()
arr = mergeSort(arr)
end1 := time.Since(start1)
//fmt.Println(arr)
fmt.Println("\n time to sort an array in normal Mergesort:\t", end1,"\n")


fmt.Println("==================== concurrent Mergesort =================")

/*here is the most important part
the function mergesortC take an array and number of cpu
and it will make go routine as many as it is efficient (base on the cpu cores that you have)
*/
start3 := time.Now()
arr3 = mergeSortC(arr3, runtime.NumCPU())
end3 := time.Since(start3)
//fmt.Println(arr3)
fmt.Println("\n time to sort an array in concurrent mode:\t", end3,"\n")

fmt.Println("==================== golang sort  =================")

start4 := time.Now()
sort.Ints(arr4)
end4 := time.Since(start4)
//fmt.Println(arr4)
fmt.Println("\n time to sort an array with golang sort builtin function :\t", end4,"\n")

}

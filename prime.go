package main

import "math"

//enter the i number it should return if it's prime or not (it is optimal)
func isPrime(i int) bool {
	for k := 2; k <= int(math.Sqrt(float64(i))); k++ {
		if i%k == 0 {
			return false
		}
	}
	return true
}



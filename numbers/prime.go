package numbers

import "math"

//IsPrime is a func that you enter the i number it should return if it's prime or not (it is optimal)
func IsPrime(i int) bool {
	for k := 2; k <= int(math.Sqrt(float64(i))); k++ {
		if i%k == 0 {
			return false
		}
	}
	return true
}

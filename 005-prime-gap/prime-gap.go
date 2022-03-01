// Gap in Primes https://www.codewars.com/kata/561e9c843a2ef5a40c0000a4/go
package main

import (
	"math"
)

func Gap(g, m, n int) []int {
	// create a new list of primes in range 'm' to 'n' inclusive
	var pir []int
	for _, prime := range SieveOfEratosthenes(n + 1) {
		if prime >= m {
			pir = append(pir, prime)
		}
	}

	if len(pir) >= 2 {
		low := pir[0]
		high := pir[1]
		for _, value := range pir {
			if isGap(low, high, g) {
				return []int{low, high}
			}

			low = high
			high = value
		}
	}
	return nil // list of valid primes is less than two therefore no gap can be calculated, or no gap was found
}

func isGap(a int, b int, c int) bool {
	if b-a == c {
		return true
	}

	return false
}

// https://www.thepolyglotdeveloper.com/2016/12/determine-number-prime-using-golang/
func SieveOfEratosthenes(value int) []int {
	f := make([]bool, value)
	var listOfPrimes = []int{}
	for i := 2; i <= int(math.Sqrt(float64(value))); i++ {
		if f[i] == false {
			for j := i * i; j < value; j += i {
				f[j] = true
			}
		}
	}
	for i := 2; i < value; i++ {
		if f[i] == false {
			listOfPrimes = append(listOfPrimes, i)
		}
	}
	return listOfPrimes
}

func main() {

}

// Gap in Primes https://www.codewars.com/kata/561e9c843a2ef5a40c0000a4/go
package main

import (
	"fmt"
	"math"
)

func Gap(g, m, n int) []int {
	var pir []int
	for _, prime := range SieveOfEratosthenes(n + 1) {
		if prime >= m {
			pir = append(pir, prime)
		}
	}

	if len(pir) < 2 {
		return nil // list of valid primes is less than two therefore no gap can be calculated, return nil
	} else if testPair(pir[0], pir[1], g) {
		return pir[0:2] // if the length is at least two, and the first two produce a match, no need to go further
	} else {
		low := pir[0]
		high := pir[1]
		for index, value := range pir {
			if index+1 > len(pir) {
				fmt.Println("reached end")
			}
			if testPair(low, high, g) {
				pir = []int{low, high}
				return pir
			} else {
				low = high
				high = value
			}
		}
	}
	return nil
}

func testPair(a int, b int, c int) bool {
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
	fmt.Println(Gap(2, 5, 7))
	fmt.Println(Gap(2, 100, 110))
	fmt.Println(Gap(4, 100, 110))
	fmt.Println(Gap(10, 300, 400))
}

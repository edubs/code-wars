package main

import (
	"fmt"
)

func main() {
	fmt.Println(BouncingBall(3, .66, 1.5))
	fmt.Println(BouncingBall(50, .66, 1.5))
	fmt.Println(BouncingBall(40, .4, 10))
	fmt.Println(BouncingBall(3, 1, 1.5))
	fmt.Println(BouncingBall(40, 1, 10))
	fmt.Println(BouncingBall(5, -1, 1.5))
}

func BouncingBall(h, bounce, window float64) int {
	if h > 0 && bounce < 1 && bounce > 0 && window < h {
		count := 0
		h *= bounce
		for h > window {
			count += 2
			h *= bounce
		}
		return count + 1
	} else {
		return -1
	}
}

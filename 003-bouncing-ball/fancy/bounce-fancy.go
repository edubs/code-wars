package main

func BouncingBall(h, bounce, window float64) int {
	if h < 0 || bounce <= 0 || 1 <= bounce || h < window {
		return -1
	}

	var count int = -1
	// different 'while' loop
	for ; h > window; h *= bounce {
		count += 2
	}

	return count
}

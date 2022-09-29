package main

import (
	"fmt"
	"math"
)

// Check if number is prime
func is_prime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	//test is_prime
	for i := 0; i < 32; i++ {
		fmt.Printf("%d: %v\n", i, is_prime(i))
	}
}
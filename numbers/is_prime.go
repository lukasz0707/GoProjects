package main

import (
	"fmt"
	"math"
)

// Check if number is prime
func isPrime(n int) bool {
	if n == 2 {
		return true
	}
	if n < 2 {
		return false
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	//test isPrime
	for i := 0; i < 32; i++ {
		fmt.Printf("%d: %v\n", i, isPrime(i))
	}
}

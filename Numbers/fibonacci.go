package main

import (
	"fmt"
	"math/big"
)

// Performant big fibonacci n-th number sequential approach
func fibonacci(n int) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}
	var p, q = big.NewInt(0), big.NewInt(1)
	for i := 1; i < n; i++ {
		p.Add(p, q)
		q, p = p, q
	}

	return q
}

// If you want the output to go into a file try:
// go run fibonnaci.go > filename.txt
func main() {
	fmt.Println(fibonacci(100))
}

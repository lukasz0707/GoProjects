package main

import (
	"fmt"
	"math"
)

func primeFactors(n int) {
	c := 2
	limit := int(math.Sqrt(float64(n)))
	for n > 1 {

		if n%c == 0 {
			fmt.Print(c, " ")
			n = n / c
		} else if c > limit {
			fmt.Println(n)
			break
		} else {
			c = c + 1
		}

	}
}

func main() {
	primeFactors(101)
}

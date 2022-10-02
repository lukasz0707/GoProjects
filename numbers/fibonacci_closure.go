package main

import "fmt"

// fibonacci closure
func fibbonaciClos() func() int {
	f1, f2 := 0, 1
	return func() int {
		f := f1
		f1, f2 = f2, f2+f
		return f
	}
}

func main() {
	fib := fibbonaciClos()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}

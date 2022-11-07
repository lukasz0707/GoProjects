package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(100)
	}
	return slice
}

func selectSort(items []int) {
	n := len(items)
	for i := 0; i < n; i++ {
		minIdx := i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
}

func main() {
	arr := generateSlice(20)
	fmt.Println("arr: ", arr)
	selectSort(arr)
	fmt.Println("arrSorted: ", arr)
}

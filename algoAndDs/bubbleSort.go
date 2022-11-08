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

func bubbleSort(arr []int) []int {
	n := len(arr)
	sorted := false

	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
	return arr
}

func main() {
	arr := generateSlice(20)
	fmt.Println("arr: ", arr)
	bubbleSort(arr)
	fmt.Println("arrSorted: ", arr)
}

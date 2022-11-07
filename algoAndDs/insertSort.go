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

func insertSort(arr []int) {
	var n = len(arr)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
				j = j - 1
			} else {
				break
			}
		}
	}
}

func main() {
	arr := generateSlice(20)
	fmt.Println("arr: ", arr)
	insertSort(arr)
	fmt.Println("arrSorted: ", arr)
}

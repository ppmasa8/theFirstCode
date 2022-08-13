package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Println(arr)

	// sorted
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}

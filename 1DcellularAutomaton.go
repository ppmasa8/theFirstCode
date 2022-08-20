package main

import "fmt"

func main() {
	n := 51
	prevState := make([]int, n+2)
	prevState[n/2+1] = 1
	state := make([]int, n+2)
	copy(state, prevState)
	// rule 90
	// https://ja.wikipedia.org/wiki/%E3%82%BB%E3%83%AB%E3%83%BB%E3%82%AA%E3%83%BC%E3%83%88%E3%83%9E%E3%83%88%E3%83%B3
	next := []int{0, 1, 0, 1, 1, 0, 1, 0}
	for i := 0; i < 30; i++ {
		for j := 1; j <= n; j++ {
			state[j] = next[prevState[j-1]<<2+prevState[j]<<1+prevState[j+1]]
		}
		fmt.Println(state)
		copy(prevState, state)
	}
}
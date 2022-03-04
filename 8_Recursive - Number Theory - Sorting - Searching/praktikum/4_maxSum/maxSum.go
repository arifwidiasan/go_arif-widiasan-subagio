package main

import (
	"fmt"
)

func MaxSequence(arr []int) int {
	// func to compare
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	max_now := arr[0]
	curr := arr[0]
	//find max sum of sequence array
	for i := 1; i < len(arr); i++ {
		curr = max(arr[i], curr+arr[i])
		max_now = max(max_now, curr)
	}
	return max_now
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))   // 7
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))   // 7
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))   // 8
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))    // 12
}

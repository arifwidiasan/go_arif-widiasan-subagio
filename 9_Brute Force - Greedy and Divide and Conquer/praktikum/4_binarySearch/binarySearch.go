package main

import (
	"fmt"
	"sort"
)

func BinarySearch(array []int, x int) {
	// your code here
	sort.Ints(array)
	hasil := -1
	low := 0
	high := len(array) - 1

	for low <= high {
		median := (low + high) / 2

		if array[median] == x {
			hasil = median
			break
		}

		if array[median] < x {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	fmt.Println(hasil)
}

func main() {
	BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3)                  // 2
	BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5)                 // 3
	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53)  // 6
	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100) // -1

}

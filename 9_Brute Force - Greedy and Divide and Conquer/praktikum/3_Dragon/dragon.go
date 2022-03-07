package main

import (
	"fmt"
	"sort"
)

func DragonOfLoowater(dragonHead, knightHeight []int) {
	// your code here
	var min int
	var mati int
	sort.Ints(dragonHead)
	sort.Ints(knightHeight)
	if len(dragonHead) > len(knightHeight) {
		min = 0
	} else {
		for i := 0; i < len(dragonHead); i++ {
			for j := 0; j < len(knightHeight); j++ {
				if dragonHead[i] <= knightHeight[j] {
					min += knightHeight[j]
					knightHeight[j] = 0
					mati += 1
					break
				}
			}
		}
	}
	if mati != len(dragonHead) {
		fmt.Println("knight fall")
	} else {
		fmt.Println(min)
	}
}

func main() {
	DragonOfLoowater([]int{5, 4}, []int{7, 8, 4})    // 11
	DragonOfLoowater([]int{5, 10}, []int{5})         // knight fall
	DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2}) // knight fall
	DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5}) // 10
}

package main

import "fmt"

func PairSum(arr []int, target int) []int {
	var temp int
	var pair []int
	mapping := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		temp = target - arr[i]
		if mapping[temp] == 1 {
			pair = append(pair, mapping[temp-1], i)
		}
		mapping[arr[i]] = 1
	}
	return pair
}

func main() {
	fmt.Println("array : [1, 2, 3, 4, 6], target = 6")
	fmt.Printf("pair sum found at index: ")
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6))
	fmt.Println("array : [2, 5, 9, 11], target = 11")
	fmt.Printf("pair sum found at index : ")
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))
}

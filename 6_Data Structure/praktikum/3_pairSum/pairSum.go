package main

import "fmt"

func PairSum(arr []int, target int) []int {
	var pair []int
	mapping := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if _, ok := mapping[target-arr[i]]; ok {
			pair = append(pair, mapping[target-arr[i]], i)
		}
		mapping[arr[i]] = i
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
	fmt.Println("array : [1, 3, 5, 7], target = 12")
	fmt.Printf("pair sum found at index : ")
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))
	fmt.Println("array : [1, 4, 6, 8], target = 10")
	fmt.Printf("pair sum found at index : ")
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))
	fmt.Println("array : [1, 5, 6, 7], target = 6")
	fmt.Printf("pair sum found at index : ")
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))
}

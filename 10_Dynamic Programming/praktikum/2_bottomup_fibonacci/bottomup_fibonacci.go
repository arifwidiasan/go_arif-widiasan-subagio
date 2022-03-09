package main

import (
	"fmt"
)

func fibo(n int) int {
	mapping := make(map[int]int)
	// your code here
	for i := 0; i <= n; i++ {
		if i <= 1 {
			mapping[i] = i
		} else {
			mapping[i] = mapping[i-1] + mapping[i-2]
		}
	}
	return mapping[n]
}

func main() {
	fmt.Println(fibo(0))  // 0
	fmt.Println(fibo(1))  // 1
	fmt.Println(fibo(2))  // 1
	fmt.Println(fibo(3))  // 2
	fmt.Println(fibo(5))  // 5
	fmt.Println(fibo(6))  // 8
	fmt.Println(fibo(7))  // 13
	fmt.Println(fibo(9))  // 34
	fmt.Println(fibo(10)) // 55
}

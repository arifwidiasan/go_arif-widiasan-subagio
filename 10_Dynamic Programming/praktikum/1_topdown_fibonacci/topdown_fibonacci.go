package main

import (
	"fmt"
)

var mapping = map[int]int{}

func fibo(n int) int {
	// your code here
	nilai, counted := mapping[n]
	if counted {
		return nilai
	}

	if n <= 1 {
		mapping[n] = n
	} else {
		mapping[n] = fibo(n-2) + fibo(n-1)
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

package main

import "fmt"

func pangkat(base, pangkat int) int {
	var sum int = base
	for i := 1; i < pangkat; i++ {
		sum = sum * base
	}
	return sum
}

func main() {
	fmt.Println(pangkat(2, 3))
	fmt.Println(pangkat(5, 3))
	fmt.Println(pangkat(10, 2))
	fmt.Println(pangkat(2, 5))
	fmt.Println(pangkat(7, 3))
}

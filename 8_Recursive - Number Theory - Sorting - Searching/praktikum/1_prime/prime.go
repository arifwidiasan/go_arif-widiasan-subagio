package main

import (
	"fmt"
)

func primeX(number int) int {
	// func to determine number is prime or not
	var prime func(int, int) bool
	prime = func(n int, i int) bool {
		if n <= 2 {
			return n == 2
		}
		if n%i == 0 {
			return false
		}
		if i*i > n {
			return true
		}
		return prime(n, i+1)
	}

	if number < 1 {
		return 0
	}

	var angka, count, i int = 1, 0, 2
	for count < number {
		angka++
		if prime(angka, i) {
			count += 1
		}
	}
	return angka
}

func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29

}

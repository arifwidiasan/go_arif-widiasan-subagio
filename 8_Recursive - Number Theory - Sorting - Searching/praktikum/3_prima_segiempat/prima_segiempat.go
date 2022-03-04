package main

import (
	"fmt"
)

func primaSegiEmpat(high, wide, start int) {
	// recursive func to determine number is prime or not
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

	var findPrime bool
	sum, mulai := 0, 2
	angka := start
	// print prime
	for i := 0; i < wide; i++ {
		for j := 0; j < high; j++ {
			findPrime = false
			for !findPrime {
				angka++
				if prime(angka, mulai) {
					sum += angka
					fmt.Print(angka, " ")
					findPrime = true
				}
			}
		}
		fmt.Println()
	}
	// print sum of all prime number
	fmt.Println(sum)
}

func main() {
	primaSegiEmpat(2, 3, 13)
	fmt.Println()
	primaSegiEmpat(5, 2, 1)

}

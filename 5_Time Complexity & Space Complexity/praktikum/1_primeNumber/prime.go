package main

import (
	"fmt"
	"math"
)

func primeNumber(number int) bool {
	for i := 2; float64(i) <= math.Sqrt(float64(number)); i++ {
		if number%i == 0 {
			return false
		}
	}
	return number > 1
}

func main() {
	fmt.Println("prime number :", primeNumber(1000000007))
	fmt.Println("prime number :", primeNumber(13))
	fmt.Println("prime number :", primeNumber(17))
	fmt.Println("prime number :", primeNumber(20))
	fmt.Println("prime number :", primeNumber(35))
}

package main

import "fmt"

func pow(x, n int) int {
	hasil := 1

	for n > 0 {
		if n%2 == 1 {
			hasil = hasil * x
		}
		n = n / 2
		x = x * x
	}
	return hasil
}

func main() {
	fmt.Println("pow(2,3) =", pow(2, 3))
	fmt.Println("pow(5,3) =", pow(5, 3))
	fmt.Println("pow(10,2) =", pow(10, 2))
	fmt.Println("pow(2,5) =", pow(2, 5))
	fmt.Println("pow(7,3) =", pow(7, 3))
}

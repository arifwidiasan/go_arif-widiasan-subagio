package main

import (
	"fmt"
)

func getMinMax(number ...*int) (min int, max int) {
	a, b := *number[0], *number[0]
	for _, value := range number {
		if *value < a {
			a = *value
		}
		if *value > b {
			b = *value
		}
	}
	return a, b
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai min = ", min)
	fmt.Println("Nilai max = ", max)
}

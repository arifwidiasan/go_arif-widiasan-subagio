package main

import "fmt"

func cetakTabelPerkalian(n int) {
	var x int
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			x = i * j
			if x <= 9 {
				fmt.Printf("  %d ", x)
			} else if x <= 99 {
				fmt.Printf(" %d ", x)
			} else {
				fmt.Printf("%d ", x)
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	cetakTabelPerkalian(9)
}

package main

import "fmt"

func playWithAsterisk(n int) {
	for a := 1; a <= n; a++ {
		for i := 1; i <= n-a; i++ {
			fmt.Printf(" ")
		}
		for j := 1; j <= a; j++ {
			fmt.Printf("*")
			fmt.Printf(" ")
		}
		fmt.Printf("\n")
	}
}

func main() {
	playWithAsterisk(5)
}

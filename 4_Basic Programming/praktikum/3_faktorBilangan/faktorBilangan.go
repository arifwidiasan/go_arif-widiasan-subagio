package main

import "fmt"

func main() {
	var bilangan int
	fmt.Printf("masukkan bilangan = ")
	fmt.Scanf("%d\n", &bilangan)

	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Printf("%d\n", i)
		}
	}
}

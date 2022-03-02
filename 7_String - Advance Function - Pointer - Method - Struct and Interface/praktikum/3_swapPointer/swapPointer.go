package main

import (
	"fmt"
)

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	a := 10
	b := 20
	fmt.Println("before swap a = ", a, " b = ", b)
	swap(&a, &b)
	fmt.Println("after swap a = ", a, " b = ", b)
}

package main

import (
	"fmt"
)

func SimpleEquations(a, b, c int) {
	// your code here
	var solusi int
	var ketemu int = 0
	x, y, z := 0, 0, 0
out:
	for i := 0; i <= a; i++ {
		x = i
		for j := 0; j <= a; j++ {
			y = j
			for k := 0; k <= a; k++ {
				z = k
				solusi = 0
				if x+y+z == a {
					solusi += 1
					if x*y*z == b {
						solusi += 1
						if (x*x)+(y*y)+(z*z) == c {
							solusi += 1
							if solusi == 3 {
								ketemu = 1
								break out
							}
						}
					}
				}
			}
		}
	}
	if ketemu == 1 {
		fmt.Println(x, y, z)
	} else {
		fmt.Println("no solution")
	}
}

func main() {
	fmt.Print("A : 1, B : 2, C : 3 = ")
	SimpleEquations(1, 2, 3) // no solution
	fmt.Println()
	fmt.Print("A : 6, B : 6, C : 14 = ")
	SimpleEquations(6, 6, 14) // 1 2 3

}

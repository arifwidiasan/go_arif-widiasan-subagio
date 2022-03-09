package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	var fs, ss float64
	cost := make([]int, len(jumps))
	// your code here
	for i := 1; i < len(jumps); i++ {
		ss = math.MaxInt32
		fs = float64(cost[i-1]) + math.Abs(float64(jumps[i])-float64(jumps[i-1]))
		if i > 1 {
			ss = float64(cost[i-2]) + math.Abs(float64(jumps[i])-float64(jumps[i-2]))
		}

		if fs < ss {
			cost[i] = int(fs)
		} else {
			cost[i] = int(ss)
		}
	}
	return cost[len(jumps)-1]
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}

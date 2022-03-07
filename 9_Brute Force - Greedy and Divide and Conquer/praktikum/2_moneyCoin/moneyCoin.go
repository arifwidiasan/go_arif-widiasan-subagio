package main

import (
	"fmt"
)

func moneyCoins(money int) []int {
	// your code here
	coin := []int{10000, 5000, 2000, 1000, 500, 200, 100, 50, 20, 10, 1}
	//coin := []int{1, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000}
	var tukar []int
	var i int = 0
	for money > 0 {
		if money >= coin[i] {
			money -= coin[i]
			tukar = append(tukar, coin[i])
		} else {
			i++
		}
	}
	return tukar
}

func main() {
	fmt.Println("Money = 123, Coin = ", moneyCoins(123))     // [100 20 1 1 1]
	fmt.Println("Money = 432, Coin = ", moneyCoins(432))     // [200 200 20 10 1 1]
	fmt.Println("Money = 543, Coin = ", moneyCoins(543))     // [500, 20, 20, 1, 1, 1]
	fmt.Println("Money = 7752, Coin = ", moneyCoins(7752))   // [5000, 2000, 500, 200, 50, 1, 1]
	fmt.Println("Money = 15321, Coin = ", moneyCoins(15321)) // [10000 5000 200 100 20 1]
}

package main

import (
	"fmt"
)

func playingDomino(cards [][]int, deck []int) interface{} {
	// sort cards by largest sum
	for i := 1; i < len(cards); i++ {
		j := i
		for j > 0 {
			if cards[j-1][0]+cards[j-1][1] < cards[j][0]+cards[j][1] {
				cards[j-1], cards[j] = cards[j], cards[j-1]
			}
			j = j - 1
		}
	}
	// check if cards in hand can play or not
	for i := 0; i < len(cards); i++ {
		if cards[i][0] == deck[0] || cards[i][1] == deck[0] || cards[i][0] == deck[1] || cards[i][1] == deck[1] {
			return cards[i]
		}
	}
	return "tutup kartu"
}

func main() {

	fmt.Println(playingDomino([][]int{{6, 5}, {3, 4}, {2, 1}, {3, 3}}, []int{4, 3}))
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 3}, {3, 4}, {2, 1}}, []int{3, 6}))
	fmt.Println(playingDomino([][]int{{6, 6}, {2, 4}, {3, 6}}, []int{5, 1}))

}

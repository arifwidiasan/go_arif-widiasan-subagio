package main

import (
	"fmt"
	"strconv"
	"strings"
)

func angkaMunculSekali(angka string) []int {
	s := strings.Split(angka, "")
	check := make(map[int]int)
	result := make([]int, 0)
	for _, val := range s {
		a, _ := strconv.Atoi(val)
		check[a] += 1
	}

	for letter := range check {
		if check[letter] == 1 {
			result = append(result, letter)
		}
	}

	return result
}

func main() {
	fmt.Println("string angka : 1234123 ")
	fmt.Println("angka muncul sekali =", angkaMunculSekali("1234123"))
}

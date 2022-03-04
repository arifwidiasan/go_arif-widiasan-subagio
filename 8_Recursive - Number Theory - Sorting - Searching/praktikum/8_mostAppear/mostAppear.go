package main

import (
	"fmt"
	"sort"
)

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	check := make(map[string]int)
	value := make([]string, 0)
	result := make([]pair, 0)
	// find duplicate and how many time it appears
	for _, val := range items {
		check[val] += 1
	}
	// sort by value map
	for v := range check {
		value = append(value, v)
	}
	sort.Slice(value, func(i, j int) bool {
		return check[value[i]] < check[value[j]]
	})
	// append to slice struct pair
	for _, item := range value {
		result = append(result, pair{item, check[item]})
	}
	return result

}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	// golang->1 ruby->2 js->4
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	// C->1 D->2 B->3 A->4
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
	// football->1 basketball->1 tenis->1

}

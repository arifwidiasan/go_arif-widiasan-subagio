package main

import (
	"fmt"
)

func ArrayMerge(arrayA []string, arrayB []string) []string {

	check := make(map[string]bool)
	array_merge := append(arrayA, arrayB...)
	result := make([]string, 0)
	for _, val := range array_merge {
		check[val] = true
	}
	for letter := range check {
		result = append(result, letter)
	}

	return result
}

func main() {
	fmt.Println("Input : [king,devil jin,akuma] , [eddie,steve,geese]")
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println("\nInput : [sergei,jin] , [jin,steve,bryan]")
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println("\nInput : [alisa,yoshimitsu] , [devil jin,yoshimitsu,alisa,law]")
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println("\nInput : [] , [devil jin,sergei]")
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println("\nInput : [hwoarang] , []")
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println("\nInput : [] , []")
	fmt.Println(ArrayMerge([]string{}, []string{}))
}

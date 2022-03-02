package main

import (
	"fmt"
)

func caesar(offset int, input string) string {
	runes := []rune(input)
	for index, char := range runes {
		char = (char-97+rune(offset))%26 + 97
		runes[index] = char
	}
	return string(runes)
}
func main() {
	fmt.Print("string(abc), offset(3) = ")
	fmt.Println(caesar(3, "abc"))
	fmt.Print("string(alta), offset(2) = ")
	fmt.Println(caesar(2, "alta"))
	fmt.Print("string(alterraacademy), offset(10) = ")
	fmt.Println(caesar(10, "alterraacademy"))
	fmt.Print("string(abcdefghijklmnopqrstuvwxyz), offset(1) = ")
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Print("string(abcdefghijklmnopqrstuvwxyz), offset(1000) = ")
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
}

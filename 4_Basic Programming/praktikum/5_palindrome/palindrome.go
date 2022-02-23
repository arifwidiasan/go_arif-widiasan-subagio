package main

import "fmt"

func palindrome(input string) bool {
	reverse_input := []rune(input)
	for i, j := 0, len(reverse_input)-1; i < j; i, j = i+1, j-1 {
		reverse_input[i], reverse_input[j] = reverse_input[j], reverse_input[i]
	}

	if input == string(reverse_input) {
		return true
	} else {
		return false
	}

}

func main() {
	fmt.Println(palindrome("civic"))
	fmt.Println(palindrome("katak"))
	fmt.Println(palindrome("kasur rusak"))
	fmt.Println(palindrome("mistar"))
	fmt.Println(palindrome("lion"))
}

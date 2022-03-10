package main

import (
	"fmt"
	"regexp"
)

func main() {
	freq := make(map[string]int)
	letter := "lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	letter = reg.ReplaceAllString(letter, "")
	huruf := make(chan string)
	clear := make(chan bool)
	// get character letter to channel
	go func(s string) {
		for _, v := range s {
			huruf <- string(v)
		}
	}(letter)
	// receive character from channel and input to map
	go func(s string) {
		for range s {
			found := <-huruf
			freq[found] += 1
		}
		clear <- true
	}(letter)
	//select
	select {
	case print := <-clear:
		if print {
			for i := range freq {
				fmt.Println(i, " -> ", freq[i])
			}
		}
	}
}

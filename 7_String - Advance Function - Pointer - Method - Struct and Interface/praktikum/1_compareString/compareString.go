package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	var hasil string
	if len(a) > len(b) {
		if strings.Contains(a, b) == true {
			hasil = b
		}
	} else {
		if strings.Contains(b, a) == true {
			hasil = a
		}
	}
	return hasil
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))
}

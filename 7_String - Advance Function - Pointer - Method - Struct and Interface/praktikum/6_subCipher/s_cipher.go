package main

import (
	"fmt"
)

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""
	runes := []rune(s.name)
	for index, char := range runes {
		if char <= 109 {
			char = ((109-char)+13)%26 + 97
		} else {
			char = (110-char)%26 + 109
		}
		runes[index] = char
	}
	nameEncode = string(runes)
	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""
	runes := []rune(s.nameEncode)
	for index, char := range runes {
		if char <= 109 {
			char = (110-char)%26 + 109
		} else {
			char = ((109-char)+13)%26 + 97
		}
		runes[index] = char
	}
	nameDecode = string(runes)
	return nameDecode
}

func main() {
	var menu int
	var a = student{}
	var c Chiper = &a
	fmt.Print("[1] Encrypt\n[2] Decrypt \n Choose your menu? ")
	fmt.Scan(&menu)
	if menu == 1 {
		fmt.Print("\nInput student name : ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput student Encode name : ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of student name " + a.nameEncode + " is : " + c.Decode())
	} else {
		fmt.Println("wrong input name menu")
	}
}

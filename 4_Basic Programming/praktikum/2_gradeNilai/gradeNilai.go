package main

import "fmt"

func main() {
	var studentScore int = 80
	var studentName string

	fmt.Printf("masukkan nama siswa = ")
	fmt.Scanf("%s \n", &studentName)
	fmt.Printf("masukkan nilai = ")
	fmt.Scanf("%d\n", &studentScore)
	if studentScore >= 80 && studentScore <= 100 {
		fmt.Printf("siswa %s mendapat nilai A", studentName)
	} else if studentScore >= 65 && studentScore <= 79 {
		fmt.Printf("siswa %s mendapat nilai B", studentName)
	} else if studentScore >= 50 && studentScore <= 64 {
		fmt.Printf("siswa %s mendapat nilai C", studentName)
	} else if studentScore >= 35 && studentScore <= 49 {
		fmt.Printf("siswa %s mendapat nilai D", studentName)
	} else if studentScore >= 0 && studentScore <= 34 {
		fmt.Printf("siswa %s mendapat nilai E", studentName)
	} else {
		fmt.Printf("nilai invalid")
	}
}

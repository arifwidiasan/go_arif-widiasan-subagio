package main

import "fmt"

func main() {
	t := 20
	r := 4
	const phi = 3.14
	luas_permukaan := 2 * phi * float64(r) * (float64(r) + float64(t))
	fmt.Println(luas_permukaan)

	fmt.Printf("masukkan tinggi = ")
	fmt.Scanf("%d\n", &t)
	fmt.Printf("masukkan jari-jari = ")
	fmt.Scanf("%d\n", &r)
	luas_permukaan = 2 * phi * float64(r) * (float64(r) + float64(t))
	fmt.Println(luas_permukaan)
}

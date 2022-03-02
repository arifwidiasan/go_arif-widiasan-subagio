package main

import (
	"fmt"
)

type Student struct {
	name  []string
	score []int
}

func (s Student) Avarage() float64 {
	var sum, a int
	for index := range s.name {
		sum += s.score[index]
		a += 1
	}
	return float64(sum) / float64(a)
}

func (s Student) min() (min int, name string) {
	a := 100
	var b string
	mapping := make(map[string]int)
	for index, val := range s.name {
		mapping[val] = s.score[index]
		if mapping[val] < a {
			a = mapping[val]
			b = val
		}
	}
	return a, b
}

func (s Student) max() (max int, name string) {
	var a int
	var b string
	mapping := make(map[string]int)
	for index, val := range s.name {
		mapping[val] = s.score[index]
		if mapping[val] > a {
			a = mapping[val]
			b = val
		}
	}
	return a, b
}

func main() {
	var a = Student{}

	for i := 1; i <= 5; i++ {
		var name string
		fmt.Printf("input %d student name : ", i)
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("input " + name + " score : ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}
	fmt.Println("\n\nAvarage score student is", a.Avarage())
	scoreMin, nameMin := a.min()
	fmt.Println("Min score student is :", nameMin, " (", scoreMin, ")")
	scoreMax, nameMax := a.max()
	fmt.Println("Max score student is :", nameMax, " (", scoreMax, ")")
}

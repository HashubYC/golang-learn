package day_2023_08

import "fmt"

func Day_2023_08_22() {
	f1()
	f2()
}

func f1() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s) // [0 0 0 0 0 1 2 3]
}

func f2() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3, 4)
	fmt.Println(s) // [1 2 3 4]
}

package day_2023_03

import "fmt"

// 知识点：可变函数。

func add(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

func Day_2023_03_07() {
	fmt.Println(add(1, 2))
	fmt.Println(add(1, 3, 7))
	// fmt.Println(add([]int{1, 2}))
	fmt.Println(add([]int{1, 3, 7}...))
}

package day_2023_06

// 知识点：可变函数。

func add(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

func Day_2023_06_27() {
	println(add(1, 2, 3, 4, 5))
	// println(add([]int{1,2,3}))
	println(add([]int{1, 2, 3}...))
}

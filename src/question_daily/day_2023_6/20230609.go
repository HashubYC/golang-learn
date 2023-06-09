package day_2023_06

// 知识点：可变参数函数。

// 可变参数是切片，切片是引用，所以func内赋值会带出来。

import "fmt"

func hello0609(num ...int) {
	fmt.Println(len(num), cap(num))
	num[0] = 18
}

func Day_2023_06_09() {
	i := []int{1, 2, 3, 4, 5}
	hello0609(i...)
	fmt.Println(i[0])
	fmt.Println(len(i), cap(i))
}

package day_2023_03

import "fmt"

// hello() 函数的参数在执行 defer 语句的时候会保存一份副本
// 在实际调用 hello() 函数时用

func hello(i int) {
	fmt.Println(i)
}

func Day_2023_03_03() {
	i := 5
	fmt.Println("defer 前")
	defer hello(i)
	fmt.Println("defer 后")
	i = i + 10
	fmt.Println("i+10 后")
}

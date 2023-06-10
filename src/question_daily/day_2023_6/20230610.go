package day_2023_06

// a 的类型是 int，b 的类型是 float，两个不同类型的数值不能相加，编译报错。
// Go语言的类型机制更加严格，没有隐式类型转换，所以不同类型的数据不能直接参与同一个运算。

import "fmt"

func Day_2023_06_10() {
	a := 5
	b := 8.1
	// fmt.Println(a + b)
	fmt.Println(a)
	fmt.Println(b)
}

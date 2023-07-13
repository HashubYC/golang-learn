package day_2023_05

import "fmt"

type MyInt1 int   // 基于类型 int 创建了新类型 MyInt1
type MyInt2 = int // 创建了 int 的类型别名 MyInt2
// Go 是强类型语言

func Day_2023_05_23() {
	var i int = 0
	// var i1 MyInt1 = i
	var i1 MyInt1 = MyInt1(i)
	var i2 MyInt2 = i
	fmt.Println(i1, i2)

	var j MyInt1 = 1
	fmt.Println(j, i2)
}

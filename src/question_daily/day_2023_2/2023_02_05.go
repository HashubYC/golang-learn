package day_2023_02

import "fmt"

// nil 只能赋值给指针、chan、func、interface、map 或 slice 类型的变量。
// error 类型，它是一种内置接口类型。看源码

func Day_2023_02_05() {
	// var x = nil
	// var x interface{} = nil
	// var x string = nil
	var x error = nil
	fmt.Println(x)
}

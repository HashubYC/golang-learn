package day_2023_05

import "fmt"

// 下面赋值正确的是

func Day_2023_05_28() {
	// var x = nil
	// var x interface{} = nil
	// var x string = nil
	var x error = nil

	fmt.Println(x)

	/**
	BD
	知识点是 nil
	nil 只能赋值给  指针、chan、func、interface、map 或 slice 类型的变量。
	*/
}

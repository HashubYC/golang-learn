package day_2023_02

import "fmt"

// 将函数 hello 赋值给变量 h，而不是函数的返回值（即不是进行函数调用），所以输出 not nil。
// 注意 Go 中函数是一等公民。
func hello() []string {
	return nil
}
func Day_2023_02_11() {
	// hello()返回值类型是[]string，是slice类型。
	// golang中两个slice类型是不可比较的，但是slice可以跟nil比较。
	h := hello
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}

	var a []string
	fmt.Printf("%T \t %v \t %v\n", a, a, a == nil) // []string         []      true

	var b = make([]string, 0)
	fmt.Printf("%T \t %v \t %v\n", b, b, b == nil) // []string         []      false
	// fmt.Printf("%v\n", a == b)                     // Invalid operation: a == b (the operator == is not defined on []string)
}

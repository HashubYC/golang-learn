package day_2023_07

// 基于类型创建的方法必须定义在同一个包内，上面的代码基于 int 类型创建了 PrintInt() 方法，由于 int 类型和方法 PrintInt() 定义在不同的包内，所以编译出错。

import "fmt"

// 解决的办法可以定义一种新的类型：
type MyInt int

// func (i int) PrintInt() {
func (i MyInt) PrintInt() {
	fmt.Println(i)
}

func Day_2023_07_27() {
	var i MyInt = 1
	i.PrintInt()
}

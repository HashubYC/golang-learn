package day_2023_03

import "fmt"

// 被调用函数里的 defer 语句在返回之前就会被执行，所以输出顺序是 F D M。

func f() {
	defer fmt.Println("D")
	fmt.Println("F")
}

func Day_2023_03_24() {
	f()
	fmt.Println("M")
}

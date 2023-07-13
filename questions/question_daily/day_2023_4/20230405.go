package day_2023_04

import "fmt"

// https://polarisxu.studygolang.com/posts/go/action/weekly-question-104/

// 注意这里不是定义一个结构体类型，而是定义一个结构体类型指针变量，即 x 是一个指针，指针类型是一个匿名结构体。
// x 的值是 nil，因为没有初始化，可以打印证实这一点。

/*
	x 是 nil，panic 很自然的。比如这样就会 panic：
	println(x.s)

	相应的，fmt.Println(x.s[99]) 也会 panic。但为什么 len(x.s[99]) 就不 panic 了呢？所以得从 len 入手一探究竟。
*/
func Day_2023_04_05() {
	var x *struct {
		s [][32]byte
	}

	println(len(x.s[99]))
	fmt.Printf("x.Type = %T; x.Value= %v\n", x, x)
}

package day_2023_03

// import "fmt"

// // 下面 A、B 两处应该填入什么代码，才能确保顺利打印出结果？

// /*
// f() 函数返回参数是指针类型，所以可以用 & 取结构体的指针；
// B 处，如果填 *f()，则 p 是 S 类型；如果填 f()，则 p 是 *S 类型，不过都可以使用 p.m 取得结构体的成员。
// */
// type S struct {
// 	m string
// }

// func f() *S {
// 	// return __ //A
// 	return &S{m: "foo"}
// }

// func Day_2023_03_28() {
// 	// p := __          //B
// 	p := f()
// 	fmt.Println(p.m) //print "foo"
// }

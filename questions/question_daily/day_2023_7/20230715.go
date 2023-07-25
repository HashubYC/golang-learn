package day_2023_07

import "fmt"

type Person1 struct {
	age int
}

// 前一题最后一行代码 person.age = 29 是修改引用对象的成员 age，
// 这题最后一行代码 person = &Person{29} 是修改引用对象本身，来看看有什么区别。

func Day_2023_07_15() {
	person := &Person1{28}

	// 1. 28
	// person.age 这一行代码跟之前含义是一样的，此时是将 28 当做 defer 函数的参数，
	// 会把 28 缓存在栈中，等到最后执行该 defer 语句的时候取出，即输出 28；
	defer fmt.Println(person.age)

	// 2. 28
	// defer 缓存的是结构体 Person{28} 的地址，这个地址指向的结构体没有被改变，最后 defer 语句后面的函数执行的时候取出仍是 28；
	defer func(p *Person1) {
		fmt.Println(p.age)
	}(person)

	// 3. 29
	// 闭包引用，person 的值已经被改变，指向结构体 Person{29}，所以输出 29.
	defer func() {
		fmt.Println(person.age)
	}()

	// 由于 defer 的执行顺序为先进后出，即 3 2 1，所以输出 29 28 28。

	person = &Person1{29}
}

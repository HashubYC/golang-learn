package day_2023_03

import "fmt"

// defer 关键字后面的函数或者方法想要执行必须先注册，return 之后的 defer 是不能注册的， 也就不能执行后面的函数或方法。
// https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-defer/

var a bool = true

func Day_2023_03_31() {
	defer func() {
		fmt.Println("1")
	}()
	if a == true {
		fmt.Println("2")
		return
	}
	defer func() {
		fmt.Println("3")
	}()
}

package day_2023_05

import "fmt"

// new([]int) 之后的 list 是一个 *[]int 类型的指针，不能对指针执行 append 操作
// 可以使用 make() 初始化之后再用。

// 同样的，map 和 channel 建议使用 make() 或字面量的方式初始化，不要用 new() 。

func Day_2023_05_18() {
	// list := new([]int)
	list := make([]int, 0)
	list = append(list, 1)
	fmt.Println(list)
}

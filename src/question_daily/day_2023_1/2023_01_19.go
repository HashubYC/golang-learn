package day_2023_01

import "fmt"

func Day_2023_01_19() {
	list := *new([]int)
	// 错误写法 list := new([]int)
	// list2 := make([]list)
	list = append(list, 1)
	fmt.Println(list)
}

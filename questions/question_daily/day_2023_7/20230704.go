package day_2023_07

import "fmt"

// 下面代码中 A B 两处应该怎么修改才能顺利编译？
func Day_2023_07_04() {
	m := make(map[string]int) //A
	// var m map[string]int //A
	m["a"] = 1
	if v, ok := m["b"]; ok { //B
		// if v := m["b"]; v != nil { //B
		fmt.Println(v)
	}
}

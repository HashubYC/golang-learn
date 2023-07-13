package day_2023_03

import "fmt"

// 在 A 处只声明了map m ,并没有分配内存空间，不能直接赋值，
// 需要使用 make()，都提倡使用 make() 或者字面量的方式直接初始化 map。

// B 处，v,k := m["b"] 当 key 为 b 的元素不存在的时候，v 会返回值类型对应的零值，k 返回 false。

func Day_2023_03_14() {
	// var m map[string]int //A
	m := make(map[string]int)
	m["a"] = 1
	// if v := m["b"]; v != nil {  //B
	if v, ok := m["b"]; ok { //B
		fmt.Println(v)
	}
}

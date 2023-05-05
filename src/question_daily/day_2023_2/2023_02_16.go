package day_2023_02

import "fmt"

type person struct {
	name string
}

// m 是一个 map，值是 nil。从 nil map 中取值不会报错，而是返回相应的零值，这里值是 int 类型，因此返回 0。
func Day_2023_02_16() {
	var m map[person]int
	var n map[person]float32
	var a map[person]string
	p := person{"Bob"}
	fmt.Println(m[p])
	fmt.Println(n[p])
	fmt.Println(a[p])
}
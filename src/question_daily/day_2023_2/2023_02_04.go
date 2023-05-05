package day_2023_02

import "fmt"

const (
	x = iota
	_
	y
	z = "zz"
	k
	p = iota
)

func Day_2023_02_04() {
	// 知识点 iota。
	// https://www.cnblogs.com/zsy/p/5370052.html
	fmt.Println(x, y, z, k, p)
}

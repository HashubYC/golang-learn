package day_2023_04

import "fmt"

// 知识点：iota 的用法、类型的 String() 方法。

/*
根据 iota 的用法推断出 South 的值是 2；
另外，如果类型定义了 String() 方法，当使用 fmt.Printf()、fmt.Print() 和 fmt.Println() 会自动使用 String() 方法，实现字符串的打印。
*/

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "2", "South", "West"}[d]
}
func Day_2023_04_10() {
	fmt.Println(South)
}

package day_2023_08

// 知识点：多重赋值。

/*
多重赋值分为两个步骤，有先后顺序：
	1. 计算等号左边的索引表达式和取址表达式，接着计算等号右边的表达式；
	2. 赋值；
所以本例，会先计算 s[i-1]，等号右边是两个表达式是常量，所以赋值运算等同于 i, s[0] = 2, "Z"。
*/
import "fmt"

func Day_2023_08_15() {
	i := 1
	s := []string{"A", "B", "C"}
	i, s[i-1] = 2, "Z"
	fmt.Printf("s: %v \n", s) // ZBC
}

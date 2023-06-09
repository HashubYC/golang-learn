package day_2023_02

import "fmt"

// 操作符 [i,j]。基于数组（切片）可以使用操作符 [i,j] 创建新的切片，新的切片 [)
// i、j 都是可选的，i 如果省略，默认是 0。j 如果省略，默认是原数组（切片）的长度。i、j 都不能超过这个长度值。

// 假如底层数组的大小为 k，截取之后获得的切片的长度和容量的计算方法：
// 长度：j-i，容量：k-i

// [i,j,k]，第三个参数 k 用来限制新切片的容量，但不能超过原数组（切片）的底层数组大小。截取获得的切片的长度和容量分别是：j-i、k-i。
func Day_2023_02_22() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[3:4:4]
	b := a[1:4]
	fmt.Println(t[0])
	fmt.Println(b)
}

package day_2023_06

// https://polarisxu.studygolang.com/posts/go/action/weekly-question-88/

// array/slice 的一些相关知识

/*
数组是这么说明的：
	数组是单一类型元素的有序序列，该单一类型称为元素类型。元素的个数被称为数组长度，并且不能为负值。
	长度是数组类型的一部分；它必须为一个可以被 int 类型的值所代表的非负常量。
	这里一个关键点就是，长度是数组的一部分，因此 [3]int 和 [4]int 是不同类型。
*/

/*
切片是针对一个底层数组的连续段的描述符，它提供了对该数组内有序序列元素的访问。
切片类型表示其元素类型的数组的所有切片的集合。
元素的数量被称为切片长度，且不能为负。
未初始化的切片的值为 nil 。
*/

/*
从 EBNF 的表示可以看出区别：
ArrayType   = "[", ArrayLength, "]", ElementType .
SliceType = "[", "]", ElementType .
也就是说，长度不是切片类型的一部分，切片长度可变。
*/

import "fmt"

func Day_2023_06_29() {
	a := []int{2: 1}
	fmt.Println(a)

	var x = []int{4: 44, 55, 66, 1: 77, 88}
	fmt.Println(x)
	println(len(x), x[2])

	// 数组初始化
	var intSet1 = [6]int{2, 4, 6}
	var intSet2 = [...]int{2, 4, 6}
	fmt.Println(intSet1)
	fmt.Println(intSet2)

	// 切片初始化
	// 当然，针对 Slice，更多时候是通过 make 创建，然后其他方式初始化，这里不展开了。
	var intSlice1 = []int{2, 4, 6}
	var intSlice2 = intSet1[:]
	fmt.Println(intSlice1)
	fmt.Println(intSlice2)
}

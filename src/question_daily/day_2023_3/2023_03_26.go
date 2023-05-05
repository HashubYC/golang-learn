package day_2023_03

import (
	"fmt"
	"reflect"
)

// 下面的两个切片声明中有什么区别？哪个更可取？

// 第一种切片声明不会分配内存，优先选择。

func Day_2023_03_26() {
	var a []int
	b := []int{}
	typeOfA := reflect.TypeOf(a)
	typeOfB := reflect.TypeOf(b)
	fmt.Println(typeOfA.Name(), typeOfA.Kind())
	fmt.Println(typeOfB.Name(), typeOfB.Kind())
}

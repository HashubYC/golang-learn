package day_2023_02

import "fmt"

// 知识点：结构体嵌套。
// 在嵌套结构体中，People 称为内部类型，Teacher 称为外部类型；通过嵌套，内部类型的属性、方法，可以为外部类型所有，就好像是外部类型自己的一样。

// t称之为方法ShowB的接收者，不是参数
// 代码设计时，通常将带星（ * ） 表示修改接收者的成员变量数据用

type People1 struct{}

func (p *People1) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}

func (p *People1) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People1
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func Day_2023_02_28() {
	t := Teacher{}
	t.ShowB()
	t.ShowA()
}

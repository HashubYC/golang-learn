package day_2023_10

import "fmt"

// Teacher 没有自己 ShowA()，所以调用内部类型 People 的同名方法
// Teacher 有自己 ShowB()

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func Day_2023_10_13() {
	t := Teacher{}
	t.ShowA()
	fmt.Println("------------")
	t.ShowB()
}

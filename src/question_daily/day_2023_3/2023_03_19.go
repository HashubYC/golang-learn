package day_2023_03

import "fmt"


// 知识点：defer、返回值。
// 注意一下，increaseA() 的返回参数是匿名，increaseB() 是具名。

// increaseA  返回值i=0的时候已经绑定到返回值里里，defer改i没用了。
// 先把返回变量r设为0，defer又把r改为1
func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}

func Day_2023_03_19() {
	fmt.Println(increaseA())
	fmt.Println(increaseB())
}

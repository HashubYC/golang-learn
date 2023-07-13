package day_2023_03

import "fmt"

// 知识点：defer、返回值。
// f1() =1，return 把r设成0，然后defer把r改为1 ； f2() =5，return 把r设成5，然后defer改的是t，不影响返回值 ；
// f3() =5，return 把r设成1，然后defer把r改为r+5，但是用的r是defer设定时的r，=0；（靠，是1，r+5的r不是外面的r）
func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	// 下面 r int 去掉 ，返回值为 6
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func Day_2023_03_21() {
	a := f1()
	b := f2()
	c := f3()
	fmt.Println(a, b, c)
}

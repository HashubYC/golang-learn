package day_2023_05

import "fmt"

// init() 函数是什么时候执行的？
// 一句话总结： import –> const –> var –> init() –> main()

func init() {
	fmt.Println("init1:", a)
}

func init() {
	fmt.Println("init2:", a)
}

var a = 10

const b = 100

func Day_2023_05_09() {
	fmt.Println("main:", a)

}

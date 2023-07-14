package day_2023_07

// F D M
// 被调用函数里的 defer 语句在返回之前就会被执行，所以输出顺序是 F D M。
import "fmt"

func f() {
	defer fmt.Println("D")
	fmt.Println("F")
}

func Day_2023_07_14() {
	f()
	fmt.Println("M")
}

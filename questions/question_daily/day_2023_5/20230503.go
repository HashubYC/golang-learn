package day_2023_05

import "fmt"

// 相当于问：可以对 map 的元素直接取地址吗？

// 即无法对 map 的 key 或 value 进行取址。
// 如果通过其他 hack 的方式，例如 unsafe.Pointer 等获取到了 key 或 value 的地址，
// 也不能长期持有，因为一旦发生扩容，key 和 value 的位置就会改变，之前保存的地址也就失效了。

func Day_2023_05_03() {
	m := make(map[string]int)
	// fmt.Println(&m["q"])
	println(m)     // 0xc00007e4b0
	fmt.Println(m) // map[]
}

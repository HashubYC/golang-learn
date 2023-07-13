package day_2023_04

import "fmt"

// map 的输出是无序的。
func Day_2023_04_03() {
	m := map[int]string{0: "zero", 1: "one"}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

package day_2023_02

import "fmt"

func Day_2023_02_01() {
	// str1 := 'abc' + '123'
	str2 := "abc" + "123"
	// str3 := '123' + "abc"
	str4 := fmt.Sprintf("abc%d", 123)
	fmt.Println(str2)
	fmt.Println(str4)
}

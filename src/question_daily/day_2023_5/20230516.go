package day_2023_05

// 考点：结构体比较

// 结构体比较规则注意1：只有相同类型的结构体才可以比较，结构体是否相同不但与属性类型个数有关，还与属性顺序相关。
// 结构体比较规则注意2：结构体是相同的，但是结构体属性中有不可以比较的类型，如 map,slice，则结构体不能用 == 比较。

// 可以使用 reflect.DeepEqual 进行比较

import (
	"fmt"
	"reflect"
)

func Day_2023_05_16() {

	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	if reflect.DeepEqual(sm1, sm2) {
		fmt.Println("sm1 == sm2")
	}

	// if sm1 == sm2 {
	// 	fmt.Println("sm1 == sm2")
	// }

	// a := map[string]string{"a": "1"}
	// b := map[string]string{"a": "1"}

	// if a == b {
	// 	fmt.Println("a == b")
	// }
}

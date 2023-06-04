package day_2023_06

// 考点：类型断言，类型断言的语法形如：i.(type)，
// 其中 i 是接口，type 是固定关键字，需要注意的是，只有接口类型才可以使用类型断言。

func GatValue() int {
	return 1
}

func GatValueInterFace() interface{} {
	return 1
}

func Day_2023_06_03() {
	// i1 := GatValue()
	i := GatValueInterFace()
	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}
}

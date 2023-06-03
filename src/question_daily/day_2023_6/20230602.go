package day_2023_06

func hello() []string {
	return nil
}

func Day_2023_06_02() {
	h := hello
	h1 := hello()
	if h == nil {
		println("h nil")
	} else {
		println("h not nil")
	}

	if h1 == nil {
		println("h1 nil")
	} else {
		println("h1 not nil")
	}
}

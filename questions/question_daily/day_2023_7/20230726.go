package day_2023_07

import "fmt"

// https://polarisxu.studygolang.com/posts/go/action/weekly-question-104/

func Day_2023_07_26() {
	var x *struct {
		s [][32]byte
	}
	fmt.Printf("x.Type = %T; x.Value= %v\n", x, x) // x.Type = *struct { s [][32]uint8 }; x.Value= <nil>
	// println(x.s) // panic
	println(len(x.s[99]))
}

package day_2023_03

import "fmt"

// 什么是闭包 https://en.wikipedia.org/wiki/Closure_(computer_programming)

func app() func(string) string {
	t := "Hi"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func Day_2023_03_22() {
	a := app()
	b := app()
	a("go")
	fmt.Println(a("go"))
	fmt.Println(b("All"))
	fmt.Println(a("All"))
}

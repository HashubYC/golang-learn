package day_2023_09

import "fmt"

func Day_2023_09_07() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}

	// s1 = append(s1, s2)
	s1 = append(s1, s2...)
	fmt.Println(s1)
}

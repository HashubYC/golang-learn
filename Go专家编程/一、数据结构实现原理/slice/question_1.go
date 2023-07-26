package slice

import "fmt"

func slice_1() {
	var array [5]int
	var slice = array[2:3]

	fmt.Println("lenth of slice: ", len(slice))    // 1
	fmt.Println("capacity of slice: ", cap(slice)) // 3
	fmt.Println(&slice[0] == &array[2])            // true

	slice[0] = 10
	fmt.Println(slice[0]) // 10
	fmt.Println(array[2]) // 10
}

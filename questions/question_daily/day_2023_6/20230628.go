package day_2023_06

// https://studygolang.com/topics/12982

import (
	"fmt"
	"time"
)

// func f1() {
// 	ch1 := make(chan int)
// 	go fmt.Println(<-ch1)
// 	ch1 <- 5
// 	time.Sleep(1 * time.Second)
// }

func f2() {
	ch1 := make(chan int)
	go func() {
		fmt.Println(<-ch1)
	}()
	ch1 <- 5
	time.Sleep(1 * time.Second)
}

func Day_2023_06_28() {
	// f1()
	f2()
}

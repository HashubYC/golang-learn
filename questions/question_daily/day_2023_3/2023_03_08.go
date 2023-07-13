package day_2023_03

import (
	"fmt"
	"time"
)

// go 语句后面的函数调用，其参数会先求值，这和普通的函数调用求值一样。函数调用之前，实参就被求值好了。

// 所以 go fmt.Println(<-ch1) 语句中的 <-ch1 是在 main goroutine 中求值的。
// 这相当于一个无缓冲的 chan，发送和接收操作都在一个 goroutine 中（main goroutine）进行，因此造成死锁。

// 更进一步，大家可以通过汇编看看上面两种方式的不同。

// 此外，defer 语句也要注意。比如下面的做法是不对的： defer recover()
/*
	正确的
	defer func() {
		recover()
	}()
*/

// fatal error: all goroutines are asleep - deadlock!
func Day_2023_03_08() {
	ch1 := make(chan int)
	// go fmt.Println(<-ch1)
	go func() {
		fmt.Println(<-ch1)
	}()
	ch1 <- 5
	time.Sleep(1 * time.Second)
}

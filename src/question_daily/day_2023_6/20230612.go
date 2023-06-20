package day_2023_06

// https://polarisxu.studygolang.com/posts/go/action/bytedance-interview-201112/

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func Day_2023_06_12() {
	total, sum := 0, 0
	for i := 1; i <= 10; i++ {
		sum += i
		go func() {
			total += i
		}()
	}
	fmt.Printf("total:%d \n sum:%d", total, sum)

	// Day_2023_06_12_1()
	// Day_2023_06_12_2()
	// Day_2023_06_12_3()
	// Day_2023_06_12_4()
	// Day_2023_06_12_5()
	// Day_2023_06_12_6()
	Day_2023_06_12_7()
}

// 考点一  i 使用的问题
// 会输出一堆 11（可能还有其他的数字），而不是期望的输出 1 到 10。
func Day_2023_06_12_1() {
	for i := 1; i <= 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(1e9)
}

// 改成如下
// 顺序输出但是乱序
func Day_2023_06_12_2() {
	for i := 1; i <= 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(1e9)
}

// 考点二  data race
// 因为存在多 goroutine 同时写 total 变量的问题，所以有数据竞争。可以加上 -race 参数验证
// go run -race main.go

// 故意把最后的 fmt 输出那一行去掉了，因为它用了 total 变量，避免它导致 data race。这引出考点三。
// 通过 -race 你验证，发现 data race 没了。

// 通过加锁的方式解决
func Day_2023_06_12_3() {
	var mutex sync.Mutex
	total, sum := 0, 0
	for i := 1; i <= 10; i++ {
		sum += i
		go func(i int) {
			mutex.Lock()
			total += i
			mutex.Unlock()
		}(i)
	}
	// fmt.Printf("\ntotal_3:%d \n sum_3:%d", total, sum)
}

// 也可以通过 atomic 包解决, （注意 total 的类型，因为 atomic.AddInt64 需要）
func Day_2023_06_12_4() {
	var total int64
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
		go func(i int) {
			atomic.AddInt64(&total, int64(i))
		}(i)
	}
	// fmt.Printf("\ntotal_4:%d \n sum_4:%d", total, sum)
}

// 经过上面两步，最终的结果还是不对的
// 考点三

// 初学 Go 应该遇到类似这样的问题，下面代码一般没有输出。原因是 main 函数先退出了，开启的 goroutine 根本没有机会执行。常见的解决办法是在最后加一个 Sleep
// Sleep 会让 main goroutine 休眠，调度器调度其他 goroutine 运行。
func Day_2023_06_12_5() {
	go func() {
		fmt.Println("Hello World!")
	}()
	time.Sleep(1e9)
}

// 回到开头的题目其实也存在这个问题，通过在 fmt 语句之前加上 Sleep，基本能得到正确的结果：
func Day_2023_06_12_6() {
	var total int64
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
		go func(i int) {
			atomic.AddInt64(&total, int64(i))
		}(i)
	}
	time.Sleep(1e9)

	fmt.Printf("\ntotal:%d\nsum %d\n", total, sum)
}

// 但如果加上 -race 发现还是有问题     WARNING: DATA RACE
// 所以，这种方式是不靠谱的，这时正确的方式是使用 sync.WaitGroup。
func Day_2023_06_12_7() {
	var wg sync.WaitGroup
	var total int64
	sum := 0
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		sum += i
		go func(i int) {
			defer wg.Done()
			atomic.AddInt64(&total, int64(i))
		}(i)
	}
	wg.Wait()

	fmt.Printf("\ntotal:%d\nsum %d\n", total, sum)
}

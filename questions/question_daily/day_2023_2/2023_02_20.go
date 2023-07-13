package day_2023_02

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 以下代码有什么问题，怎么解决？

func Day_2023_02_20() {
	total, sum := 0, 0

	for i := 1; i <= 10; i++ {
		sum += i
		// 第一个考点，下面两种方式等价
		// go func() {
		// 	total += i
		// }()
		go func(i int) {
			sum += i
		}(i)
	}
	time.Sleep(1e9)
	fmt.Printf("total: %d, sum: %d", total, sum)
}

// 考点二
// 该题的第二个考点：data race。
// 因为存在多 goroutine 同时写 total 变量的问题，所以有数据竞争。
// 可以加上 -race 参数验证：
// go run -race main.go

// 这可以通过加锁的方式解决：

func Day_2023_02_20_2() {
	var mutex sync.Mutex
	total1, sum1 := 0, 0

	for i := 1; i <= 10; i++ {
		sum1 += i
		go func(i int) {
			mutex.Lock()
			sum1 += i
			mutex.Unlock()
		}(i)
	}
	time.Sleep(1e9)
	fmt.Printf("total: %d, sum: %d", total1, sum1)
}

// 也可以通过 atomic 包解决：（注意 total 的类型，因为 atomic.AddInt64 需要）

func Day_2023_02_20_3() {
	sum2 := 0
	var total2 int64
	for i := 1; i <= 10; i++ {
		sum2 += i
		go func(i int) {
			atomic.AddInt64(&total2, int64(i))
		}(i)
	}
	time.Sleep(1e9)
	fmt.Printf("total: %d, sum: %d", total2, sum2)
}

// 考点三
/*
	Sleep 会让 main goroutine 休眠，调度器调度其他 goroutine 运行。

	回到开头的题目其实也存在这个问题，通过在 fmt 语句之前加上 Sleep，基本能得到正确的结果：

	但如果加上 -race 发现还是有问题：


	所以，这种方式是不靠谱的，这时正确的方式是使用 sync.WaitGroup。
*/
func Day_2023_02_20_4() {
	var wg sync.WaitGroup
	var total4 int64
	sum4 := 0
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		sum4 += i
		go func(i int) {
			defer wg.Done()
			atomic.AddInt64(&total4, int64(i))
		}(i)
	}
	wg.Wait()

	fmt.Printf("total:%d sum %d", total4, sum4)
}

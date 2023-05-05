package day_2023_05

// Go 的 map 可以边遍历边删除吗？
func Day_2023_05_01() {
	// map 并不是一个线程安全的数据结构。同时读写一个 map 是未定义的行为，如果被检测到，会直接 panic。

	// 一般而言，这可以通过读写锁来解决：sync.RWMutex。
	// 读之前调用 RLock() 函数，读完之后调用 RUnlock() 函数解锁；写之前调用 Lock() 函数，写完之后，调用 Unlock() 解锁。
}

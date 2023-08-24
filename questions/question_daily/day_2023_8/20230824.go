package day_2023_08

// 第二个返回值没有命名。

// func sum(x, y int)(total int, error) {
// func sum(x, y int)(total int, err error) {
func sum(x, y int) (int, error) {
	return x + y, nil
}

func Day_2023_08_24() {
	println(sum(1, 2))
}

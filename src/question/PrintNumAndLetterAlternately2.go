package question

import "fmt"

// 分别使用三个 channel 来控制数字，字母以及终止信号的输入。

func PrintNumAndLetterAlternately2() {
	number, letter, done := make(chan bool), make(chan bool), make(chan bool)

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
			}
		}
	}()

	go func() {
		j := 'A'
		for {
			select {
			case <-letter:
				if j >= 'Z' {
					done <- true
				} else {
					fmt.Print(string(j))
					j++
					fmt.Print(string(j))
					j++
					number <- true
				}
			}
		}
	}()

	number <- true
	// <-done
	for {
		select {
		case <-done:
			return
		}
	}
}

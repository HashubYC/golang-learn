package day_2023_02

import (
	"encoding/json"
	"fmt"
	"time"
)

func Day_2023_02_08() {
	// 答案解析 : https://polarisxu.studygolang.com/posts/go/action/weekly-question-embed-time/
	// 内嵌？
	t := struct {
		time.Time
		N int
	}{
		time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
		5,
	}

	m, _ := json.Marshal(t)
	fmt.Printf("%s", m)
}

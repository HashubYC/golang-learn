package day_2023_03

import (
	"encoding/json"
	"fmt"
	"time"
)

func Day_2023_03_02() {
	t := struct {
		time.Time
		N int
	}{
		time.Date(2022, 12, 20, 0, 0, 0, 0, time.UTC),
		5,
	}

	m, _ := json.Marshal(t)
	fmt.Printf("%s", m)
}

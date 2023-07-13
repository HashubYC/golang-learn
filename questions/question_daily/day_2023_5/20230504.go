package day_2023_05

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// 如何确认两个 map 是否相等？

func Day_2023_05_05() {
	var m map[string]int
	var n map[string]int

	fmt.Println(reflect.DeepEqual(m, n))
}

func CompareMap(dm1 map[string]interface{}, dm2 map[string]interface{}) bool {
	keySlice := make([]string, 0)
	data1Slice := make([]interface{}, 0)
	data2Slice := make([]interface{}, 0)
	for key, value := range dm1 {
		keySlice = append(keySlice, key)
		data1Slice = append(data1Slice, value)
	}
	for _, key := range keySlice {
		if data, ok := dm2[key]; ok {
			data2Slice = append(data2Slice, data)
		} else {
			return false
		}
	}
	data1Bytes, _ := json.Marshal(data1Slice)
	data2Bytes, _ := json.Marshal(data2Slice)
	return string(data1Bytes) == string(data2Bytes)
}

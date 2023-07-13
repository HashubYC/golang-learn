package question

import (
	"fmt"
	"strings"
)

// 给定一个string s1和一个string s2，请返回一个bool，代表两串是否重新排列后可相同。 保证两串的长度都小于等于5000。

func Regroup() {
	fmt.Println(isRegroup("1234", "4321"))
	fmt.Println(isRegroup("abcd", "dcba"))
	fmt.Println(isRegroup("abcd", "dcbaa"))
}

func isRegroup(s1, s2 string) bool {
	sl1 := len([]rune(s1))
	sl2 := len([]rune(s2))

	if sl1 > 5000 || sl2 > 5000 || sl1 != sl2 {
		return false
	}

	for _, v := range s1 {
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}
	return true
}

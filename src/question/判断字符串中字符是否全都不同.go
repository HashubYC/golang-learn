package question

import "strings"

// 请实现一个算法，确定一个字符串的所有字符【是否全都不同】。这里我们要求【不允许使用额外的存储结构】。
// 给定一个string，请返回一个bool值,true代表所有字符全都不同，false代表存在相同的字符。
// 保证字符串中的字符为【ASCII字符】。字符串的长度小于等于【3000】。

// ASCII字符字符一共有256个

func CharactersAllDif() {
	isUniqueString("afdsfdsfds")
	isUniqueString2("afdsfdsfds")
	isUniqString3("afdsfdsfds")
}

// strings.Count 判断在一个字符串中包含的另外一个字符串的数量
func isUniqueString(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}
	for _, v := range s {
		if v > 127 {
			return false
		}
		if strings.Count(s, string(v)) > 1 {
			return false
		}
	}
	return true
}

// strings.Index,  strings.LastIndex 判断指定字符串在另外一个字符串的索引未知，分别是第一次发现位置和最后发现位置。
func isUniqueString2(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}
	for k, v := range s {
		if v > 127 {
			return false
		}
		if strings.Index(s, string(v)) != k {
			return false
		}
	}
	return true
}

// 位运算判断
func isUniqString3(s string) bool {
	if len(s) == 0 || len(s) > 3000 {
		return false
	}
	// 256 个字符 256 = 64 + 64 + 64 + 64
	var mark1, mark2, mark3, mark4 uint64
	var mark *uint64
	for _, r := range s {
		n := uint64(r)
		if n < 64 {
			mark = &mark1
		} else if n < 128 {
			mark = &mark2
			n -= 64
		} else if n < 192 {
			mark = &mark3
			n -= 128
		} else {
			mark = &mark4
			n -= 192
		}
		if (*mark)&(1<<n) != 0 {
			return false
		}
		*mark = (*mark) | uint64(1<<n)
	}
	return true
}

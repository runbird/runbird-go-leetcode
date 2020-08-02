package offer

import "strings"

//剑指 Offer 05. 替换空格
//请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
//
//
//
//示例 1：
//
//输入：s = "We are happy."
//输出："We%20are%20happy."
//
//
//限制：
//
//0 <= s 的长度 <= 10000
//(leetcode)剑指 Offer 05.
//https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/
func replaceSpace(s string) string {
	var res strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			res.WriteString("%20")
		} else {
			res.WriteByte(s[i])
		}
	}
	return res.String()
}

func replaceSpace2(s string) string {
	var str string
	for _, v := range s {
		if v == ' ' {
			str += "%20"
		} else {
			str += string(v)
		}
	}
	return str
}

func replaceSpace3(s string) string {
	res := make([]rune, len(s)*3)
	index := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			res[index] = '%'
			res[index+1] = '2'
			res[index+2] = '0'
			index += 3
		} else {
			res[index] = rune(s[i])
			index++
		}
	}
	return string(res)[:index]
}

package string

import "strconv"

//415. 字符串相加
//给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
//注意：
//
//num1 和num2 的长度都小于 5100.
//num1 和num2 都只包含数字 0-9.
//num1 和num2 都不包含任何前导零。

// 你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。
// 来源：力扣（LeetCode）415
// 链接：https://leetcode-cn.com/problems/add-strings

// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 <br>
// 创建日期： 2020/8/3 <br>
//
// @author suocaiyuan
// @version V1.0

func addStrings(num1 string, num2 string) string {
	n, m, add := len(num1)-1, len(num2)-1, 0
	var res string
	for ; n >= 0 || m >= 0 || add != 0; m, n = m-1, n-1 {
		var n1, m1 int
		if n >= 0 {
			n1 = int(num1[n] - '0')
		}
		if m >= 0 {
			m1 = int(num2[m] - '0')
		}
		temp := n1 + m1 + add
		res = strconv.Itoa(temp%10) + res
		add = temp / 10
	}
	return res
}

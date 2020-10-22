package string

//给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
//
//示例 1:
//
//输入: num1 = "2", num2 = "3"
//输出: "6"
//示例 2:
//
//输入: num1 = "123", num2 = "456"
//输出: "56088"
//说明：
//
//num1 和 num2 的长度小于110。
//num1 和 num2 只包含数字 0-9。
//num1 和 num2 均不以零开头，除非是数字 0 本身。
//不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
//
//来源：力扣（LeetCode）43
//链接：https://leetcode-cn.com/problems/multiply-strings
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func multiply(num1 string, num2 string) string {
	if "0" == num1 || "0" == num2 {
		return "0"
	}
	m, n := len(num1), len(num2)
	res := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		m1 := num1[i] - '0'
		for j := n - 1; j >= 0; j-- {
			n1 := num2[j] - '0'
			temp := res[i+j+1] + int(m1*n1)
			res[i+j+1] = temp % 10
			res[i+j] += res[i+j]
		}
	}
	// ret := strings.TrimLeft(string([]byte),"0")
	var ret string
	for i, v := range res {
		if i == 0 && v == 0 {
			continue
		}
		ret += string(v + '0')
	}
	return ret
}

func main() {
	num1, num2 := "123", "456"
	print(multiply(num1, num2))
}

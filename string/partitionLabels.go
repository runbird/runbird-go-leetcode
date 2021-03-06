package string

//字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一个字母只会出现在其中的一个片段。返回一个表示每个字符串片段的长度的列表。
//
//示例 1：
//输入：S = "ababcbacadefegdehijhklij"
//输出：[9,7,8]
//解释：
//划分结果为 "ababcbaca", "defegde", "hijhklij"。
//每个字母最多出现在一个片段中。
//像 "ababcbacadefegde", "hijhklij" 的划分是错误的，因为划分的片段数较少。
//
//提示：
//S的长度在[1, 500]之间。
//S只包含小写字母 'a' 到 'z' 。
//
//来源：力扣（LeetCode）763
//链接：https://leetcode-cn.com/problems/partition-labels
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 <br>
//创建日期： 2020/10/22 <br>
//
//@author suocaiyuan
//@version V1.0

func partitionLabels(S string) []int {
	dict := [26]int{}
	for key, _ := range S {
		dict[S[key]-'a'] = key
	}
	start, end, ans := 0, 0, []int{}
	for i := 0; i < len(S); i++ {
		//	end = int(math.Max(float64(end), float64((dict[S[i]-'a']))))
		end = Max(end, dict[S[i]-'a'])
		if end == i {
			ans = append(ans, end-start+1)
			start = end + 1
		}
	}
	return ans
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

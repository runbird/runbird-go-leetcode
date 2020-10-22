package string

//567. 字符串的排列
//给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
//换句话说，第一个字符串的排列之一是第二个字符串的子串。
//
//示例1:
//输入: s1 = "ab" s2 = "eidbaooo"
//输出: True
//解释: s2 包含 s1 的排列之一 ("ba").
//
//示例2:
//输入: s1= "ab" s2 = "eidboaoo"
//输出: False
//
//注意：
//
//输入的字符串只包含小写字母
//两个字符串的长度都在 [1, 10,000] 之间
//leetcode 567
//https://leetcode-cn.com/problems/permutation-in-string/

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	dict1, dict2 := make([]int, 26), make([]int, 26)
	for i := 0; i < len(dict1); i++ {
		dict1[s1[i]-'a']++
		dict2[s2[i]-'a']++
	}

	for i := 0; i < len(s2)-len(s1); i++ {
		if matches(dict1, dict2) {
			return true
		}
		dict2[s2[i+len(s2)]-'a']++
		dict2[s2[i]-'a']--

	}
	return matches(dict1, dict2)
}

func matches(dict1 []int, dict2 []int) bool {
	for i := 0; i < 26; i++ {
		if dict1[i] != dict2[i] {
			return false
		}
	}
	return true
}

func main() {
	s1, s2 := "ab", "eidbaooo"
	print(checkInclusion(s1, s2))

}

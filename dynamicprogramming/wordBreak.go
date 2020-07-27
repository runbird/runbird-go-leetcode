package main

///**
// * 类名： WordBreak <br>
// * 描述：给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
// * 说明：
// * 拆分时可以重复使用字典中的单词。
// * 你可以假设字典中没有重复的单词。
// *
// * 示例 1：
// * 输入: s = "leetcode", wordDict = ["leet", "code"]
// * 输出: true
// * 解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
// *
// * 示例 2：
// * 输入: s = "applepenapple", wordDict = ["apple", "pen"]
// * 输出: true
// * 解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
// *      注意你可以重复使用字典中的单词。
// *
// * 示例 3：
// * 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
// * 输出: false
// *
// * 来源：力扣（LeetCode）139
// * 链接：https://leetcode-cn.com/problems/word-break
// * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 <br>
// * 创建日期： 2020/7/27 <br>

func wordBreak(s string, wordDict []string) bool {
	wordsSet := make(map[string]bool)
	for _, v := range wordDict {
		wordsSet[v] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordsSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

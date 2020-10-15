package trie

//96. 不同的二叉搜索树
//给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？
//
//示例:
//
//输入: 3
//输出: 5
//解释:
//给定 n = 3, 一共有 5 种不同结构的二叉搜索树:
//
//   1         3     3      2      1
//    \       /     /      / \      \
//     3     2     1      1   3      2
//    /     /       \                 \
//   2     1         2                 3

//(leet-code)96
//https://leetcode-cn.com/problems/unique-binary-search-trees/

func numTrees(n int) int {
	ans := make([]int, n+1)
	ans[0] = 1
	ans[1] = 1
	for i := 2; i < n+1; i++ {
		for j := 1; j <= i; j++ {
			ans[i] += ans[j-1] * ans[i-j]
		}
	}
	return ans[n]
}

func main() {
	print(numTrees(5))
}

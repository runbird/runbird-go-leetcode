package main

import "math"

//120. 三角形最小路径和
//给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
//
//相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
//
//
//
//例如，给定三角形：
//
//[
//     [2],
//    [3,4],
//   [6,5,7],
//  [4,1,8,3]
//]
//自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
//
//
//
//说明：
//
//如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。
//
//来源：力扣（LeetCode）120
//链接：https://leetcode-cn.com/problems/triangle
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n)
	//重要的一步，填充里面的数组
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		temp := triangle[i]
		for j := 1; j < i; j++ {
			dp[i][j] = int(math.Min(float64((dp[i-1][j])), float64((dp[i-1][j-1])))) + temp[j]
		}
		dp[i][i] = dp[i-1][i-1] + temp[i]
	}

	minAns := dp[n-1][0]
	for i := 1; i < len(dp[n-1]); i++ {
		minAns = int(math.Min(float64(dp[n-1][i]), float64(minAns)))
	}
	return minAns
}

func main() {
	//nums := [][]int{{2},
	//	{3, 4},
	//	{6, 5, 7},
	//	{4, 1, 8, 3}}
	nums := [][]int{
		{-1},
		{-2, -3}}

	print(minimumTotal(nums))
}

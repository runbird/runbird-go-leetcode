package main

//给定一个整数矩阵，找出最长递增路径的长度。
//
//对于每个单元格，你可以往上，下，左，右四个方向移动。 你不能在对角线方向上移动或移动到边界外（即不允许环绕）。
//
//示例 1:
//
//输入: nums =
//[
//  [9,9,4],
//  [6,6,8],
//  [2,1,1]
//]
//输出: 4
//解释: 最长递增路径为 [1, 2, 6, 9]。
//示例 2:
//
//输入: nums =
//[
//  [3,4,5],
//  [3,2,6],
//  [2,2,1]
//]
//输出: 4
//解释: 最长递增路径是 [3, 4, 5, 6]。注意不允许在对角线方向上移动。
//
//来源：力扣（LeetCode）329
//链接：https://leetcode-cn.com/problems/longest-increasing-path-in-a-matrix
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//模板操作即可 DFS+记忆化
//列举可移动方向
var (
	dirs          = [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	rows, columns int
)

func longestIncreasingPath(matrix [][]int) int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows, columns = len(matrix), len(matrix[0])
	memo := make([][]int, rows)
	for i := 0; i < rows; i++ {
		memo[i] = make([]int, columns)
	}
	ans := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			ans = max(ans, dfs(matrix, memo, i, j))
		}
	}
	return ans
}

func dfs(matrix, memo [][]int, row, column int) int {
	if memo[row][column] != 0 {
		return memo[row][column]
	}
	memo[row][column]++
	for i := 0; i < len(dirs); i++ {
		dir := dirs[i]
		newRow, newColumn := row+dir[0], column+dir[1]
		//向下一层检索 检索值递增方向
		if newRow >= 0 && newRow < rows && newColumn >= 0 && newColumn < columns && matrix[newRow][newColumn] > matrix[row][column] {
			memo[row][column] = max(memo[row][column], dfs(matrix, memo, newRow, newColumn)+1)
		}
	}
	return memo[row][column]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	matrix := [][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1}}

	//matrix := [][]int{
	//	{3, 4, 5},
	//	{3, 2, 6},
	//	{2, 2, 1}}
	print(longestIncreasingPath(matrix))
}

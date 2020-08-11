package dfs

//130. 被围绕的区域
//给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。
//
//找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
//
//示例:
//
//X X X X
//X O O X
//X X O X
//X O X X
//运行你的函数后，矩阵变为：
//
//X X X X
//X X X X
//X X X X
//X O X X
//解释:
//
//被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。
// 如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
//https://leetcode-cn.com/problems/surrounded-regions/

//DFS遍历  先将边界位置为O的元素，关联的不合规的O 全部替换为#,最后将合规的替换为X,#替换为O
func solve(board [][]byte) {
	if board == nil || len(board) == 0 {
		return
	}
	if len(board[0]) == 0 {
		return
	}
	for i, m := 0, len(board); i < m; i++ {
		for j, n := 0, len(board[0]); j < n; j++ {
			isBorad := i == 0 || j == 0 || i == m-1 || j == n-1
			if isBorad && board[i][j] == 'O' {
				dfs(board, i, j)
			}
		}
	}
	for i, m := 0, len(board); i < m; i++ {
		for j, n := 0, len(board[0]); j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] == 'X' || board[i][j] == '#' {
		return
	}
	board[i][j] = '#'
	dfs(board, i-1, j)
	dfs(board, i+1, j)
	dfs(board, i, j-1)
	dfs(board, i, j+1)
}

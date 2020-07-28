package main

//编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：
//
//每行的元素从左到右升序排列。
//每列的元素从上到下升序排列。
//示例:
//
//现有矩阵 matrix 如下：
//
//[
//  [1,   4,  7, 11, 15],
//  [2,   5,  8, 12, 19],
//  [3,   6,  9, 16, 22],
//  [10, 13, 14, 17, 24],
//  [18, 21, 23, 26, 30]
//]
//给定 target = 5，返回 true。
//
//给定 target = 20，返回 false。
//
//来源：力扣（LeetCode）240
//链接：https://leetcode-cn.com/problems/search-a-2d-matrix-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//暴力破解法
func searchMatrix2(matrix [][]int, target int) bool {
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return false
	}
	row, column := len(matrix), len(matrix[0])
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if target == matrix[i][j] {
				return true
			}
		}
	}
	return false
}

//暴力破解法 优化
func searchMatrix3(matrix [][]int, target int) bool {
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return false
	}
	row, column := len(matrix), len(matrix[0])
	if matrix[0][0] > target || matrix[row-1][column-1] < target {
		return false
	}
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if target == matrix[i][j] {
				return true
			}
		}
	}
	return false
}

// 选左上角，往右走和往下走都增大，不能选
// 选右下角，往上走和往左走都减小，不能选
// 选左下角，往右走增大，往上走减小，可选
// 选右上角，往下走增大，往左走减小，可选
func searchMatrix4(matrix [][]int, target int) bool {
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return false
	}
	rows, columns := len(matrix), len(matrix[0])
	row, column := rows-1, 0
	for row >= 0 && column < columns {
		if matrix[row][column] > target {
			row--
		} else if matrix[row][column] < target {
			column++
		} else {
			return true
		}
	}
	return false
}

package array

//编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
//
//每行中的整数从左到右按升序排列。
//每行的第一个整数大于前一行的最后一个整数。
//示例 1:
//
//输入:
//matrix = [
//  [1,   3,  5,  7],
//  [10, 11, 16, 20],
//  [23, 30, 34, 50]
//]
//target = 3
//输出: true
//示例 2:
//
//输入:
//matrix = [
//  [1,   3,  5,  7],
//  [10, 11, 16, 20],
//  [23, 30, 34, 50]
//]
//target = 13
//输出: false
//
//来源：力扣（LeetCode）74
//链接：https://leetcode-cn.com/problems/search-a-2d-matrix
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m < 1 {
		return false
	}
	n := len(matrix[0])
	left, right := 0, m*n-1
	for left <= right {
		midIndex := left + (right-left)/2
		midValue := matrix[midIndex/n][midIndex%n]
		if midValue == target {
			return true
		} else if target > midValue {
			left = midIndex + 1
		} else {
			right = midIndex - 1
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50}}
	print(searchMatrix(matrix, 11))
}

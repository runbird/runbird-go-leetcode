package array

import "sort"

//找出数组中重复的数字。
//
//
//在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
//
//示例 1：
//
//输入：
//[2, 3, 1, 0, 2, 5, 3]
//输出：2 或 3
//
//
//限制：
//
//2 <= n <= 100000
//
//来源：力扣（LeetCode）剑指 Offer 03. 数组中重复的数字
//链接：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//排序 最近对比
func findRepeatNumber1(nums []int) int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return nums[i]
		}
	}
	return -1
}

//临时表对比
func findRepeatNumber2(nums []int) int {
	n := len(nums)
	temp := make([]int, n)
	for i := 0; i < n; {
		temp[nums[i]]++
		if temp[nums[i]] > 1 {
			return nums[i]
		}
	}
	return -1
}

//临时表更进一步
func findRepeatNumber4(nums []int) int {
	for i, temp := 0, 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			temp = nums[i]
			nums[i] = nums[temp]
			nums[temp] = temp
		}
	}
	return -1
}

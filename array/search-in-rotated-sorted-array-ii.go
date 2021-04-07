package array

//已知存在一个按非降序排列的整数数组 nums ，数组中的值不必互不相同。
//在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转 ，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。
//例如， [0,1,2,4,4,4,5,6,6,7] 在下标 5 处经旋转后可能变为 [4,5,6,6,7,0,1,2,4,4] 。
//给你 旋转后 的数组 nums 和一个整数 target ，请你编写一个函数来判断给定的目标值是否存在于数组中。如果 nums 中存在这个目标值 target ，则返回 true ，否则返回 false 。
//
//示例 1：
//输入：nums = [2,5,6,0,0,1,2], target = 0
//输出：true
//
//示例 2：
//输入：nums = [2,5,6,0,0,1,2], target = 3
//输出：false
//
//来源：力扣（LeetCode）81
//链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := left + (right-left)/2
		m := nums[middle]
		if m == target {
			return true
		} else if m == nums[left] {
			//无法分辨前有序还是后有序 例如：  10111和 11101
			left++
		} else if m > nums[left] {
			//前半部分有序
			if nums[left] <= target && target < m {
				//target在前半部分
				right = middle - 1
			} else {
				left = middle + 1
			}
		} else {
			if m < target && target <= nums[right] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		}
	}
	return false
}

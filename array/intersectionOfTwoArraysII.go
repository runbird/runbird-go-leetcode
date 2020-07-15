package main

import "sort"

//350. 两个数组的交集 II
//给定两个数组，编写一个函数来计算它们的交集。
//
//示例 1：
//输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2,2]

//示例 2:
//输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出：[4,9]
//
//
//说明：
//输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。
//我们可以不考虑输出结果的顺序。

//进阶：
//如果给定的数组已经排好序呢？你将如何优化你的算法？
//如果 nums1 的大小比 nums2 小很多，哪种方法更优？
//如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	up, down, ans := 0, 0, []int{}
	for up < len(nums1) && down < len(nums2) {
		if nums1[up] == nums2[down] {
			ans = append(ans, nums1[up])
			up++
			down++
			continue
		}

		if nums1[up] > nums2[down] {
			down++
		} else {
			up++
		}
	}
	return ans
}

func intersect2(nums1 []int, nums2 []int) []int {
	ret := make([]int, 0)
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	count, dict := 0, make(map[int]int, len(nums1))
	for _, v := range nums1 {
		dict[v]++
		count++
	}

	for _, v := range nums2 {
		if dict[v] > 0 && count > 0 {
			ret = append(ret, v)
			dict[v]--
			count--
		}
	}
	return ret
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{1, 2, 3, 4}
	ans := intersect2(nums1, nums2)
	for _, v := range ans {
		print(v)
	}
}

package trie

//在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为“根”。 除了“根”之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。
//
//计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。
//
//示例 1:
//
//输入: [3,2,3,null,3,null,1]
//
//     3
//    / \
//   2   3
//    \   \
//     3   1
//
//输出: 7
//解释: 小偷一晚能够盗取的最高金额 = 3 + 3 + 1 = 7.
//示例 2:
//
//输入: [3,4,5,1,3,null,1]
//
//     3
//    / \
//   4   5
//  / \   \
// 1   3   1
//
//输出: 9
//解释: 小偷一晚能够盗取的最高金额 = 4 + 5 = 9.
//
//来源：力扣（LeetCode）337
//链接：https://leetcode-cn.com/problems/house-robber-iii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func rob(root *TreeNode) int {
	ans := robReverse(root)
	return robMax(ans[0], ans[1])
}

//当前节点选择不偷[0]：当前节点能偷到的最大钱数 = 左孩子能偷到的钱 + 右孩子能偷到的钱
//当前节点选择偷[1]：当前节点能偷到的最大钱数 = 左孩子选择自己不偷时能得到的钱 + 右孩子选择不偷时能得到的钱 + 当前节点的钱数
//作者：reals
//链接：https://leetcode-cn.com/problems/house-robber-iii/solution/san-chong-fang-fa-jie-jue-shu-xing-dong-tai-gui-hu/
//来源：力扣（LeetCode)
func robReverse(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	result := make([]int, 2)
	left, right := robReverse(root.Left), robReverse(root.Right)
	result[0] = robMax(left[0], left[1]) + robMax(right[0], right[1])
	result[1] = root.Val + left[0] + right[1]
	return result
}

func robMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

package main

//108. 将有序数组转换为二叉搜索树
//将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
//
//本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
//
//示例:
//
//给定有序数组: [-10,-3,0,5,9],
//
//一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：
//
//      0
//     / \
//   -3   9
//   /   /
// -10  5

//中序遍历不能唯一确定一棵二叉搜索树。
//先序和后序遍历不能唯一确定一棵二叉搜索树。
//先序/后序遍历和中序遍历的关系：
//inorder = sorted(postorder) = sorted(preorder)，
//中序+后序、中序+先序可以唯一确定一棵二叉树。
//
//因此，“有序数组 -> BST”有多种答案。
//
//作者：LeetCode
//链接：https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/solution/jiang-you-xu-shu-zu-zhuan-huan-wei-er-cha-sou-s-15/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func sortedArrayToBST(nums []int) *TreeNode {
	ret := &TreeNode{}

	if nums == nil {
		return ret
	}

	for i := 0; i < len(nums); i++ {
		curVal := nums[i]
		ret.Val = curVal

	}

	return ret
}

func main() {

}

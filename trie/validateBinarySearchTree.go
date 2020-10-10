package main

import "math"

/*
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
示例 1:

输入:
    2
   / \
  1   3
输出: true
示例 2:

输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。

来源：力扣（LeetCode）98
链接：https://leetcode-cn.com/problems/validate-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
var pre *TreeNode

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !isValidBST(root.Left) {
		return false
	}
	if pre != nil && pre.Val >= root.Val {
		return false
	}
	pre = root
	return isValidBST(root.Right)
}

func isValidBST2(root *TreeNode) bool {
	return validBST(root, math.MinInt64, math.MaxInt64)
}

func validBST(root *TreeNode, min, max int) bool {
	//递归结束条件
	if root == nil {
		return true
	}
	// 判断节点的值是不是在区间呢，不是的话就false结束
	if root.Val <= min || root.Val >= max {
		return false
	}
	//左递归 最大值改为当前节点值
	//右递归 最小值改为当前节点值
	return validBST(root.Left, min, root.Val) && validBST(root.Right, root.Val, max)
}

func main() {
	root := TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	print(isValidBST(&root))
}

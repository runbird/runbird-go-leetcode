package main

import "strconv"

var ans []string

func binaryTreePaths(root *TreeNode) []string {
	ans = []string{}
	dfs2(root, "")
	return ans
}

func dfs2(root *TreeNode, sub string) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		sub += strconv.Itoa(root.Val)
		ans = append(ans, sub)
		return
	}
	sub += strconv.Itoa(root.Val)
	sub += "->"
	dfs2(root.Left, sub)
	dfs2(root.Right, sub)
}

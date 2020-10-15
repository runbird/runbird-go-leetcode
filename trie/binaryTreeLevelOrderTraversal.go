package trie

import (
	"fmt"
)

//102. 二叉树的层序遍历
//给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
//
//
//
//示例：
//二叉树：[3,9,20,null,null,15,7],
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//返回其层次遍历结果：
//
//[
//  [3],
//  [9,20],
//  [15,7]
//]

func levelOrderError(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := [][]int{}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		var row []int
		for i := 0; i < len(queue); i++ {
			cur := queue[i]
			row = append(row, cur.Val)
			if cur.Left != nil {
				//queue越来越大
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		result = append(result, row)
	}
	return result
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := [][]int{}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		copyQueue := []*TreeNode{}
		var row []int
		for i := 0; i < len(queue); i++ {
			cur := queue[i]
			row = append(row, cur.Val)
			if cur.Left != nil {
				copyQueue = append(copyQueue, cur.Left)
			}
			if cur.Right != nil {
				copyQueue = append(copyQueue, cur.Right)
			}
		}
		result = append(result, row)
		queue = copyQueue
	}
	return result
}

//	作者：LeetCode-Solution
//	链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal/solution/er-cha-shu-de-ceng-xu-bian-li-by-leetcode-solution/
//	来源：力扣（LeetCode）
//	著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func demoFromLeetCode(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}

func main() {
	root := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   27,
				Left:  nil,
				Right: nil,
			},
		},
	}

	result := levelOrder(&root)
	fmt.Println(result)
}

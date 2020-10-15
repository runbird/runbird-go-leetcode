package trie

//114. 二叉树展开为链表
//给定一个二叉树，原地将它展开为一个单链表。
//
//
//
//例如，给定二叉树
//
//    1
//   / \
//  2   5
// / \   \
//3   4   6
//将其展开为：
//
//1
// \
//  2
//   \
//    3
//     \
//      4
//       \
//        5
//         \
//          6
// 来源：力扣（LeetCode）114
// 链接：https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
// @program: runbird-leetCode
// @author: Suocaiyuan
// @date: 2020-08-02 13:38

//方法一 前序遍历
func flatten(root *TreeNode) {
	list := reverse(root)
	for i := 1; i < len(list); i++ {
		pre, cur := list[i-1], list[i]
		pre.Left, pre.Right = nil, cur
	}
}

func reverse(root *TreeNode) []*TreeNode {
	list := []*TreeNode{}
	if root != nil {
		list = append(list, root)
		list = append(list, reverse(root.Left)...)
		list = append(list, reverse(root.Right)...)
	}
	return list
}

//方法二：前序遍历和展开同步进行
func flatten2(root *TreeNode) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}
	var pre *TreeNode
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre == nil {
			pre.Left = nil
			pre.Right = cur
		}
		left, right := cur.Left, cur.Right
		if right == nil {
			stack = append(stack, right)
		}
		if left == nil {
			stack = append(stack, left)
		}
		pre = cur
	}
}

//方法三：寻找前驱节点
func flatten3(root *TreeNode) {
	cur := root
	for cur != nil {
		if cur.Left != nil {
			next := cur.Left
			predecessor := next
			for predecessor.Right != nil {
				predecessor = predecessor.Right
			}
			predecessor.Right = cur.Right
			cur.Left = nil
			cur.Right = next
		}
		cur = cur.Right
	}
}

package main

//101. 对称二叉树
//给定一个二叉树，检查它是否是镜像对称的。
//
//
//
//例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
//
//    1
//   / \
//  2   2
// / \ / \
//3  4 4  3
//
//
//但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
//
//    1
//   / \
//  2   2
//   \   \
//   3    3
//
//
//进阶：
//
//你可以运用递归和迭代两种方法解决这个问题吗？
//(leet-code)101
//https://leetcode-cn.com/problems/symmetric-tree/
func isSymmetric(root *TreeNode) bool {
	return recurse(root, root)
}

func recurse(root, mirror *TreeNode) bool {
	if root == nil && mirror == nil {
		return true
	}
	if root == nil || mirror == nil {
		return false
	}
	return root.Val == mirror.Val && recurse(root.Left, root.Right) && recurse(root.Right, root.Left)
}

func isSymmetric2(root *TreeNode) bool {
	p, q := root, root
	//var queue []*TreeNode
	queue := []*TreeNode{}

	queue = append(queue, p)
	queue = append(queue, q)

	for len(queue) > 0 {
		p, q = queue[0], queue[1]
		queue = queue[2:]
		if p == nil && q == nil {
			continue
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		queue = append(queue, p.Left)
		queue = append(queue, q.Right)

		queue = append(queue, p.Right)
		queue = append(queue, q.Left)
	}
	return true
}

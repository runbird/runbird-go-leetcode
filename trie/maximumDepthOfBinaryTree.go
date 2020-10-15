package trie

//给定一个二叉树，找出其最大深度。
//
//二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
//
//说明: 叶子节点是指没有子节点的节点。
//
//示例：
//给定二叉树 [3,9,20,null,null,15,7]，
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//返回它的最大深度 3 。
//
//来源：力扣（LeetCode）104
//链接：https://leetcode-cn.com/problems/maximum-depth-of-binary-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//递归法
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		left := maxDepth(root.Left)
		right := maxDepth(root.Right)
		return max(left, right) + 1
	}
}

type node struct {
	val   *TreeNode
	depth int
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	} else {

	}
	num := 0
	var queue []*node
	queue = append(queue, &node{root, 1})
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		num = max(num, cur.depth)
		if cur.val.Left != nil {
			queue = append(queue, &node{cur.val.Left, cur.depth + 1})
		}
		if cur.val.Right != nil {
			queue = append(queue, &node{cur.val.Right, cur.depth + 1})
		}
	}
	return num
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

package offer

//剑指 Offer 34. 二叉树中和为某一值的路径
//输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。
//
//示例:
//给定如下二叉树，以及目标和 sum = 22，
//              5
//             / \
//            4   8
//           /   / \
//          11  13  4
//         /  \    / \
//        7    2  5   1
//返回:
//
//[
//   [5,4,11,2],
//   [5,8,4,5]
//]
//
//提示：
//节点总数 <= 10000
// https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/

func pathSum(root *TreeNode, sum int) [][]int {
	ans := [][]int{}
	dfs(root, sum, &ans, []int{})
	return ans
}

//		作者：sakura-151
//		链接：https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/solution/golang-di-gui-by-sakura-151-2/
//		来源：力扣（LeetCode）
//		著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func dfs(root *TreeNode, sum int, ans *[][]int, arr []int) {
	if root == nil {
		return
	}
	arr = append(arr, root.Val)
	if root.Left == nil && root.Right == nil && sum == root.Val {
		//slice是一个指向底层的数组的指针结构体
		//因为是先序遍历，如果 root.Right != nil ,arr 切片底层的数组会被修改
		//所以这里需要 copy arr 到 subList，再添加进 ans，防止 arr 底层数据修改带来的错误
		subList := make([]int, len(arr))
		copy(subList, arr)
		*ans = append(*ans, subList)
	}
	dfs(root.Left, sum-root.Val, ans, arr)
	dfs(root.Right, sum-root.Val, ans, arr)

	arr = arr[:len(arr)-1]
}

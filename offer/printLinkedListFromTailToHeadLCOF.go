package offer

//剑指 Offer 06. 从尾到头打印链表
//输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
//
//
//
//示例 1：
//
//输入：head = [1,3,2]
//输出：[2,3,1]
//
//
//限制：
//
//0 <= 链表长度 <= 10000
//https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/

func reversePrint(head *ListNode) []int {
	cur := head
	stack := []*ListNode{}
	for cur != nil {
		stack = append(stack, cur)
		cur = cur.Next
	}
	res := make([]int, len(stack))

	for i := 0; len(stack) > 0; i++ {
		res[i] = stack[len(stack)-1].Val
		stack = stack[:len(stack)-1]
	}
	return res
}

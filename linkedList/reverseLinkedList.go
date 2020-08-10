package main

//206. 反转链表
//反转一个单链表。
//
//示例:
//输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
//
//来源：力扣（LeetCode）206
//https://leetcode-cn.com/problems/reverse-linked-list/

//递归法
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return cur
}

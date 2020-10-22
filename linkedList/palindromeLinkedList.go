package linkedList

//请判断一个链表是否为回文链表。
//
//示例 1:
//
//输入: 1->2
//输出: false
//示例 2:
//
//输入: 1->2->2->1
//输出: true
//进阶：
//你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
//
//来源：力扣（LeetCode）234
//链接：https://leetcode-cn.com/problems/palindrome-linked-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func isPalindrome(head *ListNode) bool {

	fast, slow := head, head
	//找到中点 然后对后半段反转，再一一对比 fast slow
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	//如果fast不为空，说明链表的长度是奇数个
	if fast != nil {
		slow = slow.Next
	}

	//反转后半部分链表
	slow = reverseList(slow)
	fast = head
	for slow != nil {
		if fast.Val != slow.Val {
			return false
		}
		fast = fast.Next
		slow = slow.Next
	}
	return true
}

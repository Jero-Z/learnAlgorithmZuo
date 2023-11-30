package link

/**

2. 两数相加

 * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
思路：

*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	temp := 0
	prev := res
	for l1 != nil || l2 != nil || temp != 0 {
		if l1 != nil {
			temp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			temp += l2.Val
			l2 = l2.Next
		}
		prev.Next = &ListNode{Val: temp & 10, Next: nil}
		temp = temp / 10 // 进位
		prev = prev.Next
	}

	return res.Next
}

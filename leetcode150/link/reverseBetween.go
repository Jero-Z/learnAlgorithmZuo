package link

/**
92. 反转链表 II
Middle
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。
请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]
思路：
[5,4,3,2,1]
[]
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head // 指向头节点
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next // 查询到反转的开始位置[1,2,3,4,5]
	}

	cur := pre.Next
	for i := 0; i < right-left; i++ { // 进行子范围的反转
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next // 最终将left 的next 指向right位置的节点
	}
	return dummyNode.Next //返回head
}

package link

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

/**
234. 回文链表
*/
func isPalindrome(head *ListNode) bool {

	if head == nil {
		return false
	}
	if head.Next == nil {
		return true
	}

	var arrNodes []int
	var i = 0

	for {
		arrNodes = append(arrNodes, head.Val)
		i++
		if head.Next == nil {
			break
		}
		head = head.Next

	}

	mid := i / 2
	for j := 0; j < mid; j++ {
		i--
		if arrNodes[j] == arrNodes[i] {
			continue
		}

		return false
	}
	return true
}

package link

import "testing"

func TestIsPalindrome(t *testing.T) {
	// 测试回文链表
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}
	res := isPalindrome(l)
	if !res {
		t.Errorf("expected true, got %v", res)
	}

	// 测试非回文链表
	l = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 1,
			},
		},
	}
	res = isPalindrome(l)
	if res {
		t.Errorf("expected false, got %v", res)
	}
}

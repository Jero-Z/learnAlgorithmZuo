package link

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	rear := ListNode{}
	head := ListNode{
		Val:  3,
		Next: &rear,
	}
	rear.Val = 1
	rear.Next = &head
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "c1", args: args{head: &head}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

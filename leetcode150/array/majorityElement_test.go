package array

import "testing"

func Test_majorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: args{
			nums: []int{3, 2, 3},
		}, want: 3},
		{name: "test2", args: args{
			nums: []int{0, 0, 0, 1, 1, 1, 1, 0, 0},
		}, want: 0},
		{name: "test3", args: args{
			nums: []int{2, 2, 2, 2, 1, 1, 3, 3},
		}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.args.nums); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

package array

import "testing"

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{{
		name: "case 1",
		args: args{
			nums: []int{1, 2, 3, 4, 5},
			val:  2,
		},
		want: 4,
	},
		{
			name: "case 2",
			args: args{
				nums: []int{1, 3, 4, 5},
				val:  2,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

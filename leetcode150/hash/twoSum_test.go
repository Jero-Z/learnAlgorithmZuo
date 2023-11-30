package hash

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "c1", args: args{nums: []int{2, 7, 11, 15}, target: 9}, want: []int{0, 1}},
		{name: "c2", args: args{nums: []int{3, 2, 4}, target: 6}, want: []int{1, 2}},
		{name: "c3", args: args{nums: []int{3, 3}, target: 6}, want: []int{0, 1}},
		{name: "c3", args: args{nums: []int{-3, 4, 3, 90}, target: 0}, want: []int{0, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

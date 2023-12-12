package baseSort

import (
	"reflect"
	"testing"
)

func Test_mergeSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "c1", args: args{nums: []int{2, 1, 4, 3, 6, 5}}, want: []int{1, 2, 3, 4, 5, 6}},
		{name: "c2", args: args{nums: []int{1, 1, 1, 1, 1}}, want: []int{1, 1, 1, 1, 1}},
		{name: "c3", args: args{nums: []int{2}}, want: []int{2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

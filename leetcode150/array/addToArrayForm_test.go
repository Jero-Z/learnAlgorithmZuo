package array

import (
	"reflect"
	"testing"
)

func Test_addToArrayForm(t *testing.T) {
	type args struct {
		num []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "c1", args: args{num: []int{1, 2, 0, 0}, k: 34}, want: []int{1, 2, 3, 4}},
		{name: "c2", args: args{num: []int{2, 7, 4}, k: 181}, want: []int{4, 5, 5}},
		{name: "c3", args: args{num: []int{2, 1, 5}, k: 806}, want: []int{1, 0, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addToArrayForm(tt.args.num, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addToArrayForm() = %v, want %v", got, tt.want)
			}
		})
	}
}

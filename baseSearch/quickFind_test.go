package baseSearch

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "c1", args: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9}, want: 4},
		{name: "c2", args: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 2}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findLeft(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "c1", args: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9}, want: 4},
		{name: "c2", args: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 2}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLeft(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findRight(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "c1", args: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9}, want: 4},
		{name: "c2", args: args{nums: []int{-1, 0, 3, 5, 9, 12}, target: 2}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRight(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findPeakElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "c1", args: args{nums: []int{1, 2, 3, 1}}, want: 2},
		{name: "c2", args: args{nums: []int{1, 2, 1, 3, 5, 6, 4}}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPeakElement(tt.args.nums); got != tt.want {
				t.Errorf("findPeakElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

package hash

import "testing"

func Test_isAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "c1", args: args{
			s: "anagram", t: "nagaram",
		}, want: true},
		{name: "c2", args: args{
			s: "rat", t: "car",
		}, want: false},
		{name: "c3", args: args{
			s: "aa", t: "bb",
		}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}

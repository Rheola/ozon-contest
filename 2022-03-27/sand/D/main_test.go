package main

import "testing"

func Test_check(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				s: "cba",
				t: "aaz",
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				s: "bc",
				t: "ab",
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				s: "a",
				t: "a",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := check(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}

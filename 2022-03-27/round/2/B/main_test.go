package main

import "testing"

func Test_moveCount(t *testing.T) {
	type args struct {
		n int
		r int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "6 2",
			args: args{
				n: 6,
				r: 2,
			},
			want: 2,
		},
		{
			name: "3 1",
			args: args{
				n: 3,
				r: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moveCount(tt.args.n, tt.args.r); got != tt.want {
				t.Errorf("moveCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_fact(t *testing.T) {
	type args struct {
		n int
		t int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				n: 3,
				t: 6,
			},
			want: 0,
		},
		{
			name: "2",
			args: args{
				n: 4,
				t: 7,
			},
			want: 3,
		},
		{
			name: "3",
			args: args{
				n: 1,
				t: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fact(tt.args.n, tt.args.t); got != tt.want {
				t.Errorf("fact() = %v, want %v", got, tt.want)
			}
		})
	}
}

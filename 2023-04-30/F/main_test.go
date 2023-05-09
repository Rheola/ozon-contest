package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				a: []int{6, 4, 3, 1, 2, 7},
			},
			want: "LLL!!R!R!L!!",
		},
		{
			name: "2",
			args: args{
				a: []int{10, 20, 40, 50},
			},
			want: "!!!!",
		},
		{
			name: "3",
			args: args{
				a: []int{50},
			},
			want: "!",
		},
		{
			name: "4",
			args: args{
				a: []int{50, 9, 22, 5, 3, 21, 7, 16, 27, 6},
			},
			want: "LLLL!R!LLLL!RRR!LLL!LL!R!R!!!",
		},
		{
			name: "3",
			args: args{
				a: []int{7, 6, 4, 3},
			},
			want: "R!R!L!!",
		},

		{
			name: "32",
			args: args{
				a: []int{21, 7, 16, 27, 6, 50, 9, 22},
			},
			want: "LLLL!RRR!LLL!LL!R!R!!!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.a); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

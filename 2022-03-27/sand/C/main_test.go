package main

import "testing"

func Test_stepCount(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				n: 654,
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				n: 5548,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stepCount(tt.args.n); got != tt.want {
				t.Errorf("stepCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

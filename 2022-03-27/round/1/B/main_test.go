package main

import "testing"

func Test_calcMaxSum(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				prices: []int{81, 22, 31, 44, 41, 78, 98},
			},
			want: 686,
		},
		{
			name: "2",
			args: args{
				prices: []int{81, 14, 88, 2, 22},
			},
			want: 308,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcMaxSum(tt.args.prices); got != tt.want {
				t.Errorf("calcMaxSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

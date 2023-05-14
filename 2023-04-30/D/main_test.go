package main

import (
	"reflect"
	"testing"
)

func Test_compress(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				arr: []int{3, 2, 1, 0, -1, 0, 6, 6, 7},
			},
			want: []int{3, -4, 0, 0, 6, 0, 6, 1},
		},
		{
			name: "2",
			args: args{
				arr: []int{0, 6, 6, 7},
			},
			want: []int{0, 0, 6, 0, 6, 1},
		},
		{
			name: "3",
			args: args{
				arr: []int{1000},
			},
			want: []int{1000, 0},
		},
		{
			name: "3",
			args: args{
				arr: []int{1, 3, 5, 7, 9, 11, 13},
			},
			want: []int{1, 0, 3, 0, 5, 0, 7, 0, 9, 0, 11, 0, 13, 0},
		},
		{
			name: "",
			args: args{
				arr: []int{1, 2, 1},
			},
			want: []int{1, 1, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compress(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compress() = %v, want %v", got, tt.want)
			}
		})
	}
}

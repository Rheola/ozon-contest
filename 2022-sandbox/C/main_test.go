package main

import (
	"reflect"
	"testing"
)

func Test_funcName(t *testing.T) {
	type args struct {
		staffs []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				staffs: []int{2, 1, 3, 1, 1, 4},
			},
			want: [][]int{
				{
					1, 2,
				},
				{
					3, 6,
				},
				{
					4, 5,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := funcName(tt.args.staffs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("funcName() = %v, want %v", got, tt.want)
			}
		})
	}
}

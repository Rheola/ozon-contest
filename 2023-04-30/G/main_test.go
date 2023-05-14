package main

import (
	"reflect"
	"testing"
)

func Test_convert(t *testing.T) {
	type args struct {
		posMap map[int]position
		maxSum int
	}
	tests := []struct {
		name string
		args args
		want []receipt
	}{
		{
			name: "1",
			args: args{
				posMap: map[int]position{
					10: {
						art: 10,
						pos: 0,
						sum: 200,
					},
					100: {
						art: 100,
						pos: 1,
						sum: 10,
					},
				},
				maxSum: 70,
			},
			want: []receipt{
				{
					s: []position{
						{
							art: 10,
							sum: 70,
						},
					},
				},
				{
					s: []position{
						{
							art: 10,
							sum: 70,
						},
					},
				},
				{
					s: []position{
						{
							art: 10,
							sum: 60,
						},
						{
							art: 100,
							sum: 10,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.posMap, tt.args.maxSum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}

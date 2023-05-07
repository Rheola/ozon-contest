package main

import (
	"reflect"
	"testing"
)

func Test_convert(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{
				line: "otLLLrRuEe256LLLN",
			},
			want: []string{
				"route",
				"256",
			},
		},
		{
			name: "2",
			args: args{
				line: "LRLUUDE",
			},
			want: []string{},
		},
		{
			name: "3",
			args: args{
				line: "itisatest",
			},
			want: []string{
				"itisatest",
			},
		},
		{
			name: "4",
			args: args{
				line: "abNcdLLLeUfNxNx",
			},
			want: []string{
				"af",
				"x",
				"xb",
				"ecd",
			},
		},
		{
			name: "12",
			args: args{
				line: "RreEfB7D2tNUzE11BN7s9kpUg9ivhBDlEBUxEBkNvUvRLg",
			},
			want: []string{
				"kgv",
				"vxg9ivh",
				"l7s9kpz72t11",
				"ref",
			},
		},
		{
			name: "20",
			args: args{
				line: "RNaUL06iDRbUqDtahNBkhDzl9LyE1gNvLU2tLLUE",
			},
			want: []string{
				"06qi",
				"abtah",
				"2tkhzly91g",
				"v",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.want) > 0 {
				if got := convert(tt.args.line); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("convert() = %v, want %v", got, tt.want)
				}
			} else {
				got := convert(tt.args.line)
				if len(got) > 0 {
					t.Errorf("2 convert() = %v, want %v", got, tt.want)
					t.Errorf("len()   %v", len(tt.want))
				}
			}
		})
	}
}

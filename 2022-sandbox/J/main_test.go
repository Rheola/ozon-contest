package main

import "testing"

func Test_rifm(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{
			name: "1",
			args: args{
				a: "task",
				b: "flask",
			},
			wantRes: 3,
		},
		{
			name: "2",
			args: args{
				a: "decide",
				b: "code",
			},
			wantRes: 2,
		},
		{
			name: "3",
			args: args{
				a: "id",
				b: "void",
			},
			wantRes: 2,
		},
		{
			name: "3",
			args: args{
				a: "void",
				b: "id",
			},
			wantRes: 2,
		},
		{
			name: "4",
			args: args{
				a: "code",
				b: "forces",
			},
			wantRes: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := rifm(tt.args.a, tt.args.b); gotRes != tt.wantRes {
				t.Errorf("rifm() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

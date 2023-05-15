package main

import "testing"

func Test_taskToPrint(t *testing.T) {
	type args struct {
		line     string
		strCount int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				line:     "7",
				strCount: 8,
			},
			want: "1-6,8",
		},
		{
			name: "case 2",
			args: args{
				line:     "1,7,1",
				strCount: 8,
			},
			want: "2-6,8",
		},
		{
			name: "case 3",
			args: args{
				line:     "1-5,1,7-7",
				strCount: 8,
			},
			want: "6,8",
		},
		{
			name: "case 4",
			args: args{
				line:     "1-5",
				strCount: 10,
			},
			want: "6-10",
		},
		{
			name: "case 5",
			args: args{
				line:     "2-2,2,2",
				strCount: 2,
			},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := taskToPrint(tt.args.line, tt.args.strCount); got != tt.want {
				t.Errorf("taskToPrint() = %v, want %v", got, tt.want)
			}
		})
	}
}

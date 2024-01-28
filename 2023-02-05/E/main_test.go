package main

import "testing"

func Test_simplify(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s: "Boooble",
			},
			want: "Booble",
		},
		{
			name: "yyyess и  ",
			args: args{
				s: "yyessss",
			},
			want: "yyess",
		},
		{
			name: "yyyess",
			args: args{
				s: "yyyess",
			},
			want: "yyess",
		},
		{
			name: "",
			args: args{
				s: "oooops",
			},
			want: "oops",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplify(tt.args.s); got != tt.want {
				t.Errorf("simplify() = %v, want %v", got, tt.want)
			}
		})
	}
}

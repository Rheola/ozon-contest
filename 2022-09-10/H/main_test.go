package main

import "testing"

func Test_validate(t *testing.T) {
	type args struct {
		color [][]string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid",
			args: args{
				color: [][]string{
					{
						"R", ".", "R", ".", "R", ".", "G",
					},
					{
						".", "Y", ".", "G", ".", "G", ".",
					},
					{
						"B", ".", "Y", ".", "V", ".", "V",
					},
				},
			},
			want: true,
		},
		{
			name: "invalid",
			args: args{
				color: [][]string{
					{
						"Y", ".", "R", ".", "B", ".", "B", ".",
					},
					{
						".", "R", ".", "R", ".", "B", ".", "V",
					},
					{
						"B", ".", "R", ".", "B", ".", "R", ".",
					},
					{
						".", "B", ".", "B", ".", "R", ".", "R",
					},
				},
			},
			want: false,
		},
	}

	//	R.R.R.G
	//	.Y.G.G.
	//	B.Y.V.V
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validate(tt.args.color); got != tt.want {
				t.Errorf("validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

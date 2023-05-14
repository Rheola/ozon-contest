package main

import "testing"

func Test_pl(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				word: "cat",
			},
			want: "cats",
		},
		{
			name: "",
			args: args{
				word: "puppy",
			},
			want: "puppies",
		},
		{
			name: "",
			args: args{
				word: "bus",
			},
			want: "buses",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pl(tt.args.word); got != tt.want {
				t.Errorf("pl() = %v, want %v", got, tt.want)
			}
		})
	}
}

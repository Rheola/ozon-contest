package main

import "testing"

func Test_calcSizes(t *testing.T) {
	type args struct {
		count int
	}
	tests := []struct {
		name      string
		args      args
		wantN1    int
		wantN24   int
		wantN2416 int
	}{
		{
			name: "385",
			args: args{
				count: 385,
			},
			wantN1:    1,
			wantN24:   0,
			wantN2416: 1,
		},
		{
			name: "23",
			args: args{
				count: 23,
			},
			wantN1:    0,
			wantN24:   1,
			wantN2416: 0,
		},
		{
			name: "27",
			args: args{
				count: 27,
			},
			wantN1:    3,
			wantN24:   1,
			wantN2416: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN1, gotN24, gotN2416 := calcSizes(tt.args.count)
			if gotN1 != tt.wantN1 {
				t.Errorf("calcSizes() gotN1 = %v, want %v", gotN1, tt.wantN1)
			}
			if gotN24 != tt.wantN24 {
				t.Errorf("calcSizes() gotN24 = %v, want %v", gotN24, tt.wantN24)
			}
			if gotN2416 != tt.wantN2416 {
				t.Errorf("calcSizes() gotN2416 = %v, want %v", gotN2416, tt.wantN2416)
			}
		})
	}
}

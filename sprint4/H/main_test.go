package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"true",
			args{
				"mxyskaoghi",
				"qodfrgmslc",
			},
			true,
		},
		{
			"true2",
			args{
				"agg",
				"xdd",
			},
			true,
		},
		{
			"true3",
			args{
				"agg",
				"xda",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

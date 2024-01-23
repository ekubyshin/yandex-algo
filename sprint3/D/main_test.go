package main

import (
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		children string
		cookies  string
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{
			"0",
			args{
				"1 2",
				"2 1 3",
			},
			2,
		},
		{
			"1",
			args{
				"2 1 3",
				"1 1",
			},
			1,
		},
		{
			"2",
			args{
				"4",
				"1 4 7 10 2 2 7 8",
			},
			1,
		},
		{
			"3",
			args{
				"8 5 5 8 6 9 8 2 4 7",
				"9 8 1 1 1 5 10 8",
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := solve(readArr(tt.args.children, len(strings.Split(tt.args.children, " "))), readArr(tt.args.cookies, len(strings.Split(tt.args.cookies, " ")))); gotRes != tt.wantRes {
				t.Errorf("solve() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

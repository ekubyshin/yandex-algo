package main

import "testing"

func Test_is_correct_bracket_seq(t *testing.T) {
	type args struct {
		arr []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"ok",
			args{
				arr: []byte("{[()]}"),
			},
			true,
		},
		{
			"ok",
			args{
				arr: []byte(""),
			},
			true,
		},
		{
			"ok",
			args{
				arr: []byte("()"),
			},
			true,
		},
		{
			"ok",
			args{
				arr: []byte("()()"),
			},
			true,
		},
		{
			"ok",
			args{
				arr: []byte("[()()]"),
			},
			true,
		},
		{
			"ok",
			args{
				arr: []byte("()())"),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := is_correct_bracket_seq(tt.args.arr); got != tt.want {
				t.Errorf("is_correct_bracket_seq() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_handleInput(t *testing.T) {
	type args struct {
		inp string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				"2 1 + 3 *",
			},
			9,
		},
		{
			"2",
			args{
				"7 2 + 4 * 2 +",
			},
			38,
		},
		{
			"3",
			args{
				"4 13 5 / +",
			},
			6,
		},
		{
			"4",
			args{
				"4 2 * 4 / 25 *",
			},
			50,
		},
		{
			"3",
			args{
				"4 2 * 4 / 25 * 2 - 12 / 500 2 * + 2 / -999 + 71 + -1 * 2 / 1000 + 6 * 8065 - 787 +",
			},
			0,
		},
		{
			"3",
			args{
				"4 2 * 4 / 25 * 2 - 12 / 500 2 * + 2 / -999 + 71 + -1 * 2 / 1000 + 6 * 8065 - 787 + 66 *",
			},
			0,
		},
		{
			"3",
			args{
				"2 5 - 4 /",
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewStack()
			arr := strings.Split(tt.args.inp, " ")
			for _, inp := range arr {
				handleInput(inp, &q)
			}
			got, _ := q.Pop()
			assert.Equal(t, tt.want, got)
		})
	}
}

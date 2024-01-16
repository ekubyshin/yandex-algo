package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solver(t *testing.T) {
	tests := []struct {
		name string
		nums []string
		want string
	}{
		{
			"1",
			strings.Split("15 56 2", " "),
			"56215",
		},
		{
			"2",
			strings.Split("1 783 2", " "),
			"78321",
		},
		{
			"3",
			strings.Split("2 4 5 2 10", " "),
			"542210",
		},
		{
			"4",
			strings.Split("9 10 1 1 1 6", " "),
			"9611110",
		},
		{
			"5",
			strings.Split("9 18 1 65 51 16 6 43 6 36 7 11 64 5 1 76 15 11 11 15 57 95 3 49 92 78 83 51 10 3", " "),
			"995928378776666564575515149433633181615151111111110",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, solver(tt.nums))
		})
	}
}

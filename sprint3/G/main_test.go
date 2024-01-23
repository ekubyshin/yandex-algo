package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solve(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			"0",
			[]int{0, 2, 1, 2, 0, 0, 1},
			[]int{0, 0, 0, 1, 1, 2, 2},
		},
		{
			"1",
			[]int{2, 1, 2, 0, 1},
			[]int{0, 1, 1, 2, 2},
		},
		{
			"2",
			[]int{2, 1, 1, 2, 0, 2},
			[]int{0, 1, 1, 2, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, solve(tt.arr))
		})
	}
}

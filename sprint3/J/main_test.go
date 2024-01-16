package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_bubbleSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			"1",
			[]int{2, 1, 3},
			[]int{1, 2, 3},
		},
		{
			"2",
			[]int{4, 5},
			[]int{4, 5},
		},
		{
			"3",
			[]int{1, 1, 1, 1, 1},
			[]int{1, 1, 1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bubbleSort(tt.arr)
			assert.Equal(t, tt.want, tt.arr)
		})
	}
}

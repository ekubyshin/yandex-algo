package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_quicksort(t *testing.T) {
	tests := []struct {
		name string
		args []Competitor
		want []Competitor
	}{
		{
			"0",
			[]Competitor{
				{
					[]byte("alla"),
					4,
					100,
				},
				{
					[]byte("gena"),
					6,
					1000,
				},
				{
					[]byte("gosha"),
					2,
					90,
				},
				{
					[]byte("rita"),
					2,
					90,
				},
				{
					[]byte("timofey"),
					4,
					80,
				},
			},
			[]Competitor{
				{
					[]byte("gena"),
					6,
					1000,
				},
				{
					[]byte("timofey"),
					4,
					80,
				},
				{
					[]byte("alla"),
					4,
					100,
				},
				{
					[]byte("gosha"),
					2,
					90,
				},
				{
					[]byte("rita"),
					2,
					90,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quicksort(tt.args, Compare)
			assert.Equal(t, tt.want, tt.args)
		})
	}
}

func Test_quicksortint(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			"0",
			[]int{
				1, 0, 3, 10, 5, 7, 6,
			},
			[]int{
				0, 1, 3, 5, 6, 7, 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quicksort(tt.args, func(a int, b int) int {
				if a > b {
					return 1
				}
				if a == b {
					return 0
				}
				return -1
			})
			assert.Equal(t, tt.want, tt.args)
		})
	}
}

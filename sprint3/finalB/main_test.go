package main

import (
	"fmt"
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
					"alla",
					4,
					100,
				},
				{
					"gena",
					6,
					1000,
				},
				{
					"gosha",
					2,
					90,
				},
				{
					"rita",
					2,
					90,
				},
				{
					"timofey",
					4,
					80,
				},
			},
			[]Competitor{
				{
					"gena",
					6,
					1000,
				},
				{
					"timofey",
					4,
					80,
				},
				{
					"alla",
					4,
					100,
				},
				{
					"gosha",
					2,
					90,
				},
				{
					"rita",
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

func Test_readCompetitors(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			"a",
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := read()
			assert.Len(t, r, tt.want)
			for _, l := range r {
				fmt.Println(l)
			}
		})
	}
}

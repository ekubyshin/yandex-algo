package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solve(t *testing.T) {
	tests := []struct {
		name  string
		pairs []Pair
		want  []Pair
	}{
		{
			"1",
			[]Pair{
				{
					7, 8,
				},
				{
					7, 8,
				},
				{
					2, 3,
				},
				{
					6, 10,
				},
			},
			[]Pair{
				{
					2, 3,
				},
				{
					6, 10,
				},
			},
		},
		{
			"2",
			[]Pair{
				{
					2, 3,
				},
				{
					5, 6,
				},
				{
					3, 4,
				},
				{
					3, 4,
				},
			},
			[]Pair{
				{
					2, 4,
				},
				{
					5, 6,
				},
			},
		},
		{
			"3",
			[]Pair{
				{
					1, 3,
				},
				{
					3, 5,
				},
				{
					4, 6,
				},
				{
					5, 6,
				},
				{
					2, 4,
				},
				{
					7, 10,
				},
			},
			[]Pair{
				{
					1, 6,
				},
				{
					7, 10,
				},
			},
		},
		{
			"4",
			[]Pair{
				{1, 10},
			},
			[]Pair{
				{1, 10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, solve(tt.pairs))
		})
	}
}

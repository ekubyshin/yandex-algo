package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_merge(t *testing.T) {
	type args struct {
		arr   []int
		left  int
		mid   int
		right int
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{
			"1",
			args{
				[]int{3, 4, 5, 0, 1, 2},
				0,
				3,
				6,
			},
			[]int{0, 1, 2, 3, 4, 5},
		},
		{
			"2",
			args{
				[]int{5, 7, 8, 3, 4, 5, 0, 1, 2},
				3,
				6,
				9,
			},
			[]int{0, 1, 2, 3, 4, 5},
		},
		{
			"3",
			args{
				[]int{5, 7, 8, 3, 4, 5, 0, 1, 2, 0, 0, 0},
				3,
				6,
				9,
			},
			[]int{0, 1, 2, 3, 4, 5},
		},
		{
			"4",
			args{
				[]int{1, 4, 9, 2, 10, 11},
				0,
				3,
				6,
			},
			[]int{1, 2, 4, 9, 10, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantResult, merge(tt.args.arr, tt.args.left, tt.args.mid, tt.args.right))
		})
	}
}

func Test_merge_sort(t *testing.T) {
	type args struct {
		arr []int
		lf  int
		rg  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"1",
			args{
				[]int{1, 4, 2, 10, 1, 2},
				0,
				6,
			},
			[]int{1, 1, 2, 2, 4, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge_sort(tt.args.arr, tt.args.lf, tt.args.rg)
			assert.Equal(t, tt.want, tt.args.arr)
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func Test_solver(t *testing.T) {
	type args struct {
		n   int
		arr []int
		pos []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test 1",
			args{
				5,
				[]int{0, 1, 4, 9, 0},
				[]int{0, 4},
			},
			[]int{0, 1, 2, 1, 0},
		},
		{
			"test 2",
			args{
				6,
				[]int{0, 1, 4, 9, 10, 0},
				[]int{0, 5},
			},
			[]int{0, 1, 2, 2, 1, 0},
		},
		{
			"test 3",
			args{
				6,
				[]int{0, 1, 4, 9, 10, 20},
				[]int{0},
			},
			[]int{0, 1, 2, 3, 4, 5},
		},
		{
			"test 4",
			args{
				3,
				[]int{0, 0, 0},
				[]int{0, 1, 2},
			},
			[]int{0, 0, 0},
		},
		{
			"test 5",
			args{
				3,
				[]int{0, 0, 1},
				[]int{0, 0},
			},
			[]int{0, 0, 1},
		},
		{
			"test 6",
			args{
				3,
				[]int{1, 0, 0},
				[]int{1, 2},
			},
			[]int{1, 0, 0},
		},
		{
			"test 7",
			args{
				9,
				[]int{98, 0, 10, 77, 0, 59, 28, 0, 94},
				[]int{1, 4, 7},
			},
			[]int{1, 0, 1, 1, 0, 1, 1, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solver(tt.args.n, tt.args.arr, tt.args.pos); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solver() = %v, want %v", got, tt.want)
			}
		})
	}
}

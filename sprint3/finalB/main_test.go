package main

import (
	"bufio"
	"bytes"
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

func Test_readCompetitors(t *testing.T) {
	type args struct {
		reader *bufio.Reader
		num    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"a",
			args{
				bufio.NewReader(bytes.NewReader([]byte(`alla 4 100
				gena 6 1000
				gosha 2 90
				rita 2 90
				timofey 4 80`))),
				5,
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := readCompetitors(tt.args.reader, tt.args.num)
			assert.Len(t, r, tt.want)
			for _, l := range r {
				fmt.Println(l)
			}
		})
	}
}

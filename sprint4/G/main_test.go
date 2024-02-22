package main

import (
	"strconv"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	tests := []struct {
		name     string
		sequence []int
		want     int
	}{
		{
			"1",
			[]int{0, 0, 1, 0, 1, 1, 1, 0, 0, 0},
			8,
		},
		{
			"2",
			[]int{0, 1},
			2,
		},
		{
			"3",
			[]int{0, 1, 0},
			2,
		},
		{
			"4",
			[]int{0, 1, 0, 1, 1, 1, 0, 0, 0},
			8,
		},
		{
			"5",
			readS("1 1 1 1 0 1 1 0 0 0 0 0 1 1 0 0 1 1 0 0 1 0 0 1 0 0 1 1 0 0 1 0 0 0 0 1 1 1 0 0 0 1 0 1 1 1 0 1 1 0 0 1 0 0 0 1 0 1 0 0 0 1 1 1 1 1 1 0 1 0 0 0"),
			46,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.sequence); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func readS(str string) []int {
	arr := strings.Split(str, " ")
	sequence := make([]int, 0, len(arr))
	for _, line := range arr {
		result, _ := strconv.Atoi(line)
		sequence = append(sequence, result)
	}
	return sequence
}

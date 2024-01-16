package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generate(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{
			"23",
			strings.Split("ad ae af bd be bf cd ce cf", " "),
		},
		{
			"92",
			strings.Split("wa wb wc xa xb xc ya yb yc za zb zc", " "),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, generate(tt.name))
		})
	}
}

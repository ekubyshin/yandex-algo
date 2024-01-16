package main

import (
	"strings"
	"testing"
)

func Test_isSequence(t *testing.T) {
	type args struct {
		source []byte
		input  []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{
				source: stringToByteArray("abc"),
				input:  stringToByteArray("ahbgdcu"),
			},
			true,
		},
		{
			"1",
			args{
				source: stringToByteArray("abcp"),
				input:  stringToByteArray("ahpc"),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSequence(tt.args.source, tt.args.input); got != tt.want {
				t.Errorf("isSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func stringToByteArray(s string) []byte {
	arr := strings.Split(s, "")
	res := make([]byte, 0, len(arr))
	for _, c := range arr {
		res = append(res, c[0])
	}
	return res
}

package main

import "testing"

func Test_hash(t *testing.T) {
	type args struct {
		str string
		a   uint64
		m   uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			"1",
			args{
				"a",
				123,
				100003,
			},
			97,
		},
		{
			"2",
			args{
				"hash",
				123,
				100003,
			},
			6080,
		},
		{
			"3",
			args{
				"HaSH",
				123,
				100003,
			},
			56156,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hash(tt.args.str, tt.args.a, tt.args.m); got != tt.want {
				t.Errorf("hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

package L

import "testing"

func Test_solver(t *testing.T) {
	type args struct {
		p     int64
		elems []int64
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			"",
			args{
				elems: []int64{1, 2, 4, 4, 6, 8},
				p:     3,
			},
			3,
			5,
		},
		{
			"",
			args{
				elems: []int64{1, 2, 4, 4, 4, 4},
				p:     3,
			},
			3,
			-1,
		},
		{
			"",
			args{
				elems: []int64{1, 2, 4, 4, 4, 4},
				p:     10,
			},
			-1,
			-1,
		},
		{
			"",
			args{
				elems: []int64{1, 2, 4, 4, 4, 4},
				p:     1,
			},
			1,
			2,
		},
		{
			"",
			args{
				elems: []int64{1, 2, 4, 4, 4, 4},
				p:     4,
			},
			3,
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := solv(tt.args.p, tt.args.elems)
			if got != tt.want {
				t.Errorf("solver() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("solver() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

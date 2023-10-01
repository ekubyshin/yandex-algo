package main

import "testing"

func TestSolution(t *testing.T) {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}
	type args struct {
		head *ListNode
		elem string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"first",
			args{
				head: &node0,
				elem: "node0",
			},
			0,
		},
		{
			"last",
			args{
				head: &node0,
				elem: "node3",
			},
			3,
		},
		{
			"medium",
			args{
				head: &node0,
				elem: "node2",
			},
			2,
		},
		{
			"not found",
			args{
				head: &node0,
				elem: "not",
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.head, tt.args.elem); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	node3 := ListNode{"node3", nil, nil}
	node2 := ListNode{"node2", &node3, nil}
	node1 := ListNode{"node1", &node2, nil}
	node0 := ListNode{"node0", &node1, nil}
	node3.prev = &node2
	node2.prev = &node1
	node1.prev = &node0
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test",
			args{
				head: &node0,
			},
			"node3node2node1node0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := make([]string, 0)
			got := Solution(tt.args.head)
			cur := got
			for cur != nil {
				p = append(p, cur.data)
				cur = cur.next
			}
			s1 := strings.Join(p, "")
			if s1 != tt.want || got.next == nil || got.prev != nil {
				t.Errorf("error want %v got %v", tt.want, s1)
			}
		})
	}
}

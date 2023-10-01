package main

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		head *ListNode
		idx  int
	}
	node5 := ListNode{"node3", nil}
	node4 := ListNode{"node2", &node5}
	node3 := ListNode{"node2", &node4}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}
	tests := []struct {
		name     string
		args     args
		want     *ListNode
		toDelete *ListNode
	}{
		{
			name: "remove first",
			args: args{
				head: &node0,
				idx:  0,
			},
			want:     &node1,
			toDelete: &node0,
		},
		{
			name: "remove last",
			args: args{
				head: &node0,
				idx:  5,
			},
			want:     &node0,
			toDelete: &node5,
		},
		{
			name: "remove second",
			args: args{
				head: &node0,
				idx:  1,
			},
			want:     &node0,
			toDelete: &node1,
		},
		{
			name: "remove pre last",
			args: args{
				head: &node0,
				idx:  4,
			},
			want:     &node0,
			toDelete: &node4,
		},
		{
			name: "remove median",
			args: args{
				head: &node0,
				idx:  2,
			},
			want:     &node0,
			toDelete: &node2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var head *ListNode
			if head = Solution(tt.args.head, tt.args.idx); !reflect.DeepEqual(head, tt.want) {
				t.Errorf("Solution() = %v, want %v", head, tt.want)
			}
			for head != nil {
				if reflect.DeepEqual(head, tt.toDelete) {
					t.Errorf("node wasnt deleted")
				}
				head = head.next
			}
		})
	}
}

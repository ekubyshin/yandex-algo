package main

// <template>
type ListNode struct {
	data string
	next *ListNode
}

// <template>

func Solution(head *ListNode, idx int) *ListNode {
	// Your code
	// ヽ(´▽`)/
	var newHead *ListNode
	if idx == 0 {
		newHead = head.next
		head.next = nil
		return newHead
	}
	newHead = head
	cur := head
	c := 0
	var prev *ListNode
	for cur != nil {
		if c == idx-1 {
			prev = cur
		} else if c == idx+1 {
			prev.next = cur
			break
		} else if prev != nil {
			prev.next = nil
		}
		cur = cur.next
		c++
	}

	return newHead
}

func test() {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}
	Solution(&node0, 2)
}

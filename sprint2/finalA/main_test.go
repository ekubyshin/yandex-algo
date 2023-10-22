package finalA

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue_PushFront(t *testing.T) {
	q := NewQueue(4)
	for i := 0; i < 4; i++ {
		q.PushFront(i)
	}
	assert.ElementsMatch(t, []int{0, 1, 2, 3}, q.queue)
	for i := 3; i >= 0; i-- {
		x, _ := q.PopFront()
		assert.Equal(t, i, x)
	}
	for i := 0; i < 4; i++ {
		q.PushBack(i)
	}
	assert.ElementsMatch(t, []int{0, 1, 2, 3}, q.queue)
	for i := 3; i >= 0; i-- {
		x, _ := q.PopBack()
		assert.Equal(t, i, x)
	}
}

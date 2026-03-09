package easy

import (
	q "dsa/structs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hasCycle(t *testing.T) {

	t.Run("", func(t *testing.T) {
		t.Parallel()
		node4 := &q.ListNode{Val: -4}
		node3 := &q.ListNode{Val: 0, Next: node4}
		node2 := &q.ListNode{Val: 2, Next: node3}
		head := &q.ListNode{Val: 3, Next: node2}
		node4.Next = node2 // создаем цикл: -4 -> 2

		assert.Equalf(t, true, hasCycle(head), "hasCycle()")
	})

}

package medium

import (
	. "leetcode/structs"
)

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var cur, prev *ListNode
	// 1 итерацию вручную шоб голову перезаписать

	// prev.Next = cur
	prev, cur = nil, head

	head = cur.Next

	cur.Next, cur.Next.Next = cur.Next.Next, cur
	prev = cur

	for cur.Next != nil && cur.Next.Next != nil {
		cur = cur.Next
		prev.Next, cur.Next, cur.Next.Next = cur.Next, cur.Next.Next, cur
		prev = cur
	}

	return head

}

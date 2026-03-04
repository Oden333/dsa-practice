package medium

import (
	. "dsa/structs"
)

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var reverseList_recursive func(q *ListNode, l int) *ListNode
	var level int = 1

	var leftNode *ListNode
	var rightNode *ListNode
	cur := head

	for level < left && cur.Next != nil {
		if level == left-1 {
			leftNode = cur
		}

		cur = cur.Next
		level++
	}

	reverseList_recursive = func(cur *ListNode, lev int) *ListNode {
		if cur == nil || cur.Next == nil || lev >= right {
			if cur.Next != nil {
				rightNode = cur.Next
			}
			return cur
		}

		newHead := reverseList_recursive(cur.Next, lev+1)

		cur.Next.Next = cur
		cur.Next = nil

		if lev == left {
			cur.Next = rightNode
			return newHead
		}

		return newHead
	}

	newHead := reverseList_recursive(cur, level)
	if leftNode != nil {
		leftNode.Next = newHead
	}

	if left == 1 {
		head = newHead
	}

	return head
}

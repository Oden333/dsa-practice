package structs

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	current := head

	for _, val := range vals[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return head
}

func (l *ListNode) String() string {
	if l == nil {
		return ""
	}
	var sb strings.Builder

	for l != nil {
		sb.WriteString(strconv.Itoa(l.Val))

		if l.Next != nil {
			sb.WriteString(" ")
		}
		l = l.Next
	}

	return sb.String()
}

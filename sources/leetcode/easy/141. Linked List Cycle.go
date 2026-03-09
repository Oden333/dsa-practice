package easy

import . "dsa/structs"

func hasCycle(head *ListNode) bool {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return false
	}
	for head != slow {
		head = head.Next
		slow = slow.Next
	}
	return true
}

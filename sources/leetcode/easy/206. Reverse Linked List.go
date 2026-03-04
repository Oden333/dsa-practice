package easy

import . "dsa/structs"

func reverseList(cur *ListNode) *ListNode {
	if cur == nil {
		return nil
	}

	var (
		prev *ListNode
	)

	for cur != nil {
		tmp := cur.Next

		cur.Next = prev
		prev = cur
		cur = tmp
	}

	return prev
}

func reverseList_recursive(cur *ListNode) *ListNode {
	// Базовый случай: пустой список или последний элемент
	if cur == nil || cur.Next == nil {
		return cur
	}

	// Рекурсивно доходим до новой головы и просто прокидываем её наверх для ответа
	newHead := reverseList_recursive(cur.Next)

	// Меняем направление текущего узла
	cur.Next.Next = cur
	cur.Next = nil

	return newHead
}

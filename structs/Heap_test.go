package structs

import (
	"container/heap"
	"fmt"
	"testing"
)

// TestMinHeap_Pop проверяет извлечение минимального
func TestMinHeap_Pop(t *testing.T) {
	h := &MinHeap{1, 2, 5, 3}
	heap.Init(h)

	got := heap.Pop(h).(int)
	if got != 1 {
		t.Errorf("Pop() = %d, want 1", got)
	}
	if got := h.Len(); got != 3 {
		t.Errorf("Len() после Pop = %d, want 3", got)
	}
	// После Pop новый минимальный должен быть на вершине
	if got := (*h)[0]; got != 2 {
		t.Errorf("h[0] после Pop = %d, want 2", got)
	}
}

// TestMinHeap_PushPopOrder проверяет, что элементы выходят по возрастанию
func TestMinHeap_PushPopOrder(t *testing.T) {
	input := []int{5, 3, 1, 7, 2, 9}
	expected := []int{1, 2, 3, 5, 7, 9}

	h := &MinHeap{}
	heap.Init(h)
	for _, v := range input {
		heap.Push(h, v)
	}

	for _, want := range expected {
		got := heap.Pop(h).(int)
		if got != want {
			t.Errorf("Pop() = %d, want %d", got, want)
		}
	}
}

func TestMinHeap_1(t *testing.T) {
	// input := []int{5, 3, 1, 7, 2, 9}
	// expected := []int{1, 2, 3, 5, 7, 9}

	h := &MinHeap{}
	for _, v := range []int{5, 3, 1, 7, 2} {
		h.Push(v)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(h.Pop()) // Должно вывести: 1 2 3 5 7
	}
}

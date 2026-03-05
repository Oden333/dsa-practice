package medium

import "container/heap"

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	h := new(ksp)
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			heap.Push(h, [2]int{v1, v2})
		}
	}

	res := make([][]int, k)
	for i := range k {
		tmp, _ := heap.Pop(h).([2]int)
		res[i] = tmp[:]
	}

	return res
}

// min heap
type ksp [][2]int

func (h ksp) Len() int {
	return len(h)
}

func (h ksp) Less(i int, j int) bool {
	return h[i][0]+h[i][1] < h[j][0]+h[j][1]

	// switch {
	// case (h[i][0] < h[j][0]):
	// 	return true
	// case (h[i][0] == h[j][0]):
	// 	return (h[i][1] <= h[j][1])
	// default:
	// 	return false
	// }
}

func (h ksp) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ksp) Push(x any) {
	(*h) = append((*h), x.([2]int))
}

func (h *ksp) Pop() any {
	val := (*h)[h.Len()-1]
	(*h) = (*h)[:h.Len()-1]
	return val
}

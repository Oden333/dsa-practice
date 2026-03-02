package structs

type monotonicStack []int

func (s *monotonicStack) Pop() int {
	idx := len(*s) - 1
	val := (*s)[idx]
	(*s) = (*s)[:idx]

	return val
}

func (s *monotonicStack) Push(val int) {
	(*s) = append((*s), val)
}

func (s *monotonicStack) Peek() int {
	return (*s)[len(*s)-1]
}

func (s *monotonicStack) Len() int {
	return len(*s)
}

func nextGreaterElement(nums []int) []int {
	s := monotonicStack{}

	for j := len(nums) - 1; j >= 0; j-- {
		val := nums[j]

		for s.Len() > 0 {
			if peekVal := s.Peek(); peekVal > val {
				nums[j] = peekVal
				break
			} else {
				_ = s.Pop()
			}
		}

		if s.Len() == 0 {
			nums[j] = -1
		}

		s.Push(val)
	}

	return nums
}

func nextLowerElement(nums []int) []int {
	s := monotonicStack{}

	for j := len(nums) - 1; j >= 0; j-- {
		val := nums[j]

		for s.Len() > 0 {
			if peekVal := s.Peek(); peekVal < val {
				nums[j] = peekVal
				break
			} else {
				_ = s.Pop()
			}
		}

		if s.Len() == 0 {
			nums[j] = -1
		}

		s.Push(val)
	}

	return nums
}

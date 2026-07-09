package scholar

type MinStack struct {
	metastack []int
	stack     []int
}

func Constructor() MinStack {
	return MinStack{
		stack:     make([]int, 0),
		metastack: make([]int, 0),
	}
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) Push(value int) {
	s.stack = append(s.stack, value)

	if len(s.metastack) == 0 || value <= s.metastack[len(s.metastack)-1] {
		s.metastack = append(s.metastack, value)
	}
}

func (s *MinStack) Pop() {
	val := s.stack[len(s.stack)-1]
	metaval := s.metastack[len(s.metastack)-1]

	if val == metaval {
		s.metastack = s.metastack[:len(s.metastack)-1]
	}

	s.stack = s.stack[:len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	return s.metastack[len(s.metastack)-1]
}

func (s *MinStack) Show() ([]int, []int) {
	return s.stack, s.metastack
}

package stack

type Stack struct {
	slice []interface{}
}

func (s *Stack) Push(v interface{}) {
	s.slice = append(s.slice, v)
}

func (s *Stack) Top() interface{} {
	if len(s.slice) == 0 {
		return nil
	}
	return s.slice[len(s.slice)-1]
}

func (s *Stack) Pop() interface{} {
	top := s.Top()
	if top == nil {
		return nil
	}
	s.slice = s.slice[:len(s.slice)-1]
	return top
}

func (s *Stack) Empty() bool {
	return len(s.slice) == 0
}

func (s *Stack) Slice() []interface{} {
	return s.slice
}

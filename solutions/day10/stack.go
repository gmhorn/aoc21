package day10

import "errors"

type Stack struct {
	elem []rune
	top  int
}

func NewStack() *Stack {
	return &Stack{
		elem: make([]rune, 0),
		top:  -1,
	}
}

func (s *Stack) Push(val rune) {
	s.elem = append(s.elem, val)
	s.top++
}

func (s *Stack) Pop() (rune, error) {
	if s.top < 0 {
		return 0, errors.New("empty stack")
	}
	val := s.elem[s.top]
	s.elem = s.elem[:s.top]
	s.top--
	return val, nil
}

func (s *Stack) Empty() bool {
	return len(s.elem) == 0
}

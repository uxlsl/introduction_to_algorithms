package main

import "fmt"

type Stack struct {
	top int
	S   []int
}

func (s *Stack) stack_empty() bool {
	if len(s.S) == 0 {
		return true
	} else {
		return false
	}
}

func (s *Stack) stack_push(v int) {
	fmt.Println(s.S, len(s.S), cap(s.S))
	if s.top < len(s.S)-1 {
		s.top += 1
		s.S[s.top] = v
	} else {
		s.S = append(s.S, v)
		s.top += 1
	}
}

func (s *Stack) stack_pop() int {
	if s.top > 0 {
		s.top--
		return s.S[s.top+1]
	}
	return -1
}

func NewStack() *Stack {
	return &Stack{-1, make([]int, 2)}
}

func main() {
	s := NewStack()
	s.stack_push(1)
	s.stack_push(2)
	s.stack_push(3)
	s.stack_push(4)
	s.stack_push(5)
	s.stack_push(5)
	s.stack_push(5)
	s.stack_push(5)
	s.stack_push(5)
	s.stack_push(5)
	fmt.Println(s.S)
}

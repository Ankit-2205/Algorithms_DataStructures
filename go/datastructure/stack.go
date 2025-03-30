package datastructure

import (
	"container/list"
)

type Stack struct {
	*list.List
}

func NewStack() Stack {
	return Stack{list.New()}
}

func (s *Stack) Push(value int) {
	s.PushBack(value)
}

func (s *Stack) Pop() int {
	e := s.Back()
	s.Remove(e)
	return e.Value.(int)
}

func (s *Stack) Top() int {
	return s.Back().Value.(int)
}

func (s *Stack) Empty() bool {
	return s.Len() == 0
}

// container/list methods for stack
// init: list.New()
// push: list.PushBack()
// top: list.Back()
// pop: list.Remove(list.Back())
// print: for e := list.Back(); e != nil; e = e.Prev() { println(e.Value) }

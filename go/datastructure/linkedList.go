package datastructure

import (
	"container/list"
)

type LinkedList struct {
	*list.List
}

func NewLinkedList() LinkedList {
	return LinkedList{list.New()}
}

func (ll *LinkedList) PushTail(value int) *list.Element {
	return ll.PushBack(value)
}

func (ll *LinkedList) PushHead(value int) *list.Element {
	return ll.PushFront(value)
}

func (ll *LinkedList) Head() int {
	return ll.Front().Value.(int)
}

func (ll *LinkedList) Tail() int {
	return ll.Back().Value.(int)
}

func (ll *LinkedList) PopBack() int {
	e := ll.Back()
	ll.Remove(e)
	return e.Value.(int)
}

func (ll *LinkedList) PopFront() int {
	e := ll.Front()
	ll.Remove(e)
	return e.Value.(int)
}

func (ll *LinkedList) Empty() bool {
	return ll.Len() == 0
}

// container/list methods for linked list
// init: list.New()
// element: *Element
// push back: list.PushBack()
// push front: list.PushFront()
// head: list.Front()
// tail: list.Back()
// pop back: list.Remove(list.Back())
// pop front: list.Remove(list.Front())
// print: for e := list.Front(); e != nil; e = e.Next() { println(e.Value) }

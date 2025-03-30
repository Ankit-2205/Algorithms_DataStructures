package datastructure

import "container/list"

type Queue interface {
	Insert(key int, value string)
	PrintQueue()
	Top() (key int, value string)
	Pop() (key int, value string)
}

// Simple Queue using container/list

type SimpleQueue struct {
	*list.List
}

func NewSimpleQueue() *SimpleQueue {
	return &SimpleQueue{list.New()}
}

func (q *SimpleQueue) Insert(key int, value string) {
	q.PushBack(Node{key, value})
}

func (q *SimpleQueue) PrintQueue() {
	for e := q.Front(); e != nil; e = e.Next() {
		node := e.Value.(Node)
		println(node.key, node.value)
	}
}

func (q *SimpleQueue) Top() (key int, value string) {
	node := q.Front().Value.(Node)
	return node.key, node.value
}

func (q *SimpleQueue) Pop() (key int, value string) {
	node := q.Front().Value.(Node)
	q.Remove(q.Front())
	return node.key, node.value
}

/* container/list methods for queue
init: list.New()
insert: list.PushBack()
top: list.Front()
pop: list.Remove(list.Front())
print: for e := list.Front(); e != nil; e = e.Next() { println(e.Value) }
*/

package datastructure

import "fmt"

type Node struct {
	key   int
	value string
}

type PriorityQueue []Node

func (pq *PriorityQueue) fixUp(index int) {

	if index == 1 {
		return
	}
	heap := *pq

	parent := index / 2
	if heap[parent].key > heap[index].key {
		t := heap[index]
		heap[index] = heap[parent]
		heap[parent] = t
	}

	pq.fixUp(parent)
}

func (pq *PriorityQueue) fixDown(index int) {

	heap := *pq
	leftChild := 2 * index
	rightChild := 2*index + 1
	minIndex := index

	if rightChild < len(heap) && heap[minIndex].key > heap[rightChild].key {
		minIndex = rightChild
	}

	if leftChild < len(heap) && heap[minIndex].key > heap[leftChild].key {
		minIndex = leftChild
	}

	if minIndex != index {
		heap[minIndex], heap[index] = heap[index], heap[minIndex]
		heap.fixDown(minIndex)
	}
}

func (pq *PriorityQueue) Top() (key int, value string) {
	return (*pq)[1].key, (*pq)[1].value
}

func (pq *PriorityQueue) Pop() (key int, value string) {
	heap := *pq

	last := len(heap) - 1
	t := heap[1]
	heap[1] = heap[last]
	heap[last] = t

	*pq = (*pq)[:last]
	pq.fixDown(1)
	return t.key, t.value
}

func (pq *PriorityQueue) Insert(key int, value string) {
	*pq = append(*pq, Node{
		key,
		value,
	})

	pq.fixUp(len(*pq) - 1)
}

func (pq *PriorityQueue) PrintQueue() {
	for _, node := range (*pq)[1:] {
		fmt.Printf("%d: %s, ", node.key, node.value)
	}

	fmt.Println()
}

func NewPriorityQueue() Queue {
	return &PriorityQueue{Node{}}
}

// Time Complexity: O(log(n))
// Space Complexity: O(n)

package aoc

import (
	"container/heap"
)

type PriorityQueue[K, T comparable] []*Item[K, T]

type Item[K, T comparable] struct {
	node     *Node[K, T]
	priority int
	index    int
}

func (pq PriorityQueue[K, T]) Len() int { return len(pq) }

func (pq PriorityQueue[K, T]) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue[K, T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue[K, T]) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item[K, T])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[K, T]) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue[K, T]) update(item *Item[K, T], node *Node[K, T], priority int) {
	item.node = node
	item.priority = priority
	heap.Fix(pq, item.index)
}

package prioqueue

import (
	"container/heap"
)

type Prioqueue[T any] struct {
	inner *pq[T]
}

func New[T any]() *Prioqueue[T] {
	return &Prioqueue[T]{inner: new(pq[T])}
}

func (p *Prioqueue[T]) Push(item T, prio int) {
	n := newNode(item, prio)
	heap.Push(p.inner, n)
}

func (p *Prioqueue[T]) PopValue() T {
	return p.PopItem().Value
}

func (p *Prioqueue[T]) PopItem() Item[T] {
	item := heap.Pop(p.inner)
	return *item.(*Item[T])
}

func (p *Prioqueue[T]) Len() int { return p.inner.Len() }

// inner implementation
type Item[T any] struct {
	Value    T
	priority int
	index    int
}

func (i Item[T]) Priority() int {
	return i.priority
}

func newNode[T any](val T, prio int) *Item[T] {
	return &Item[T]{
		Value:    val,
		priority: prio,
	}
}

type pq[T any] []*Item[T]

func (p pq[T]) Len() int {
	return len(p)
}

func (p pq[T]) Less(i, j int) bool {
	a, b := p[i], p[j]
	return a.priority < b.priority
}

func (p pq[T]) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index = i
	p[j].index = j
}

func (p *pq[T]) Push(x any) {
	n := len(*p)
	item := x.(*Item[T])
	item.index = n
	*p = append(*p, item)
}

func (p *pq[T]) Pop() any {
	old := *p
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*p = old[0 : n-1]
	return item
}

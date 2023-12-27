package prioqueue

import (
	"container/heap"
)

type Prioqueue[T comparable] struct {
	inner pq[T]
}

func New[T comparable]() *Prioqueue[T] {
	return &Prioqueue[T]{inner: pq[T]{}}
}

func (p *Prioqueue[T]) Push(item T, prio int) {
	n := newNode(item, prio)
	heap.Push(&p.inner, n)
}

func (p *Prioqueue[T]) Pop() T {
	return p.popItem().Value
}

func (p *Prioqueue[T]) IndexFunc(fn func(T) bool) int {
	for idx := range p.inner {
		if fn(p.inner[idx].Value) {
			return idx
		}
	}
	return -1
}

func (p *Prioqueue[T]) PeekItem(idx int) T {
	return p.inner[idx].Value
}

func (p *Prioqueue[T]) PriorityOfItem(idx int) int {
	return p.inner[idx].priority
}

func (p *Prioqueue[T]) UpdateAtIndex(idx int, item T, prio int) {
	p.inner[idx].Value = item
	p.inner[idx].priority = prio
	heap.Fix(&p.inner, idx)
}

func (p *Prioqueue[T]) popItem() item[T] {
	r := heap.Pop(&p.inner)
	return *r.(*item[T])
}

func (p *Prioqueue[T]) Len() int { return p.inner.Len() }

// inner implementation
type item[T comparable] struct {
	Value    T
	priority int
	index    int
}

func (i item[T]) Priority() int {
	return i.priority
}

func newNode[T comparable](val T, prio int) *item[T] {
	return &item[T]{
		Value:    val,
		priority: prio,
	}
}

type pq[T comparable] []*item[T]

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
	item := x.(*item[T])
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

package prioqueue

import (
	"container/heap"
)

type Prioqueue struct {
	inner *pq
}

func New() *Prioqueue {
	inner := make(pq, 0)
	heap.Init(&inner)
	return &Prioqueue{inner: &inner}
}

func (p *Prioqueue) Push(item string, prio int) {
	n := newNode(item, prio)
	heap.Push(p.inner, n)
}

func (p *Prioqueue) PopValue() string {
	return p.PopItem().Value
}

func (p *Prioqueue) PopItem() Item {
	item := heap.Pop(p.inner)
	return *item.(*Item)
}

func (p *Prioqueue) Len() int { return p.inner.Len() }

// inner implementation
type Item struct {
	Value    string
	priority int
	index    int
}

func (i Item) Priority() int {
	return i.priority
}

func newNode(val string, prio int) *Item {
	return &Item{
		Value:    val,
		priority: prio,
	}
}

type pq []*Item

func (p pq) Len() int {
	return len(p)
}

func (p pq) Less(i, j int) bool {
	a, b := p[i], p[j]
	return a.priority < b.priority
}

func (p pq) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index = i
	p[j].index = j
}

func (p *pq) Push(x any) {
	n := len(*p)
	item := x.(*Item)
	item.index = n
	*p = append(*p, item)
}

func (p *pq) Pop() any {
	old := *p
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*p = old[0 : n-1]

	return item
}

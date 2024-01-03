package queue

type Queue[T any] struct {
	elem []T
}

func New[T any](items ...T) Queue[T] {
	return Queue[T]{elem: items}
}

func (q *Queue[T]) Push(items ...T) {
	q.elem = append(q.elem, items...)
}

func (q *Queue[T]) Pop() T {
	item := q.elem[0]
	q.elem = q.elem[1:]
	return item
}

func (q *Queue[T]) Len() int {
	return len(q.elem)
}

func (q *Queue[T]) Empty() bool {
	return q.Len() == 0
}

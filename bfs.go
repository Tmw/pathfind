package pathfind

import (
	"slices"

	"github.com/tmw/pathfind/pkg/queue"
	"github.com/tmw/pathfind/pkg/slice"
)

type BFS[T comparable] struct {
	candidates queue.Queue[candidate[T]]
	visited    map[T]struct{}
	adapter    Adapter[T]
}

func NewBFS[T comparable](start T, adapter Adapter[T]) BFS[T] {
	sc := candidate[T]{
		coord:  start,
		parent: nil,
	}

	candidates := queue.New[candidate[T]](sc)

	return BFS[T]{
		candidates: candidates,
		visited:    make(map[T]struct{}),
		adapter:    adapter,
	}
}

func (w *BFS[T]) IsVisited(c T) bool {
	_, found := w.visited[c]
	return found
}

func (w *BFS[T]) visit(c T) {
	w.visited[c] = struct{}{}
}

func (w *BFS[T]) Walk() []T {
	for w.candidates.Len() > 0 {
		c := w.candidates.Pop()

		if w.adapter.IsFinish(c.coord) {
			return backtrace[T](c)
		}

		w.visit(c.coord)

		neighbours := w.adapter.Neighbours(c.coord)
		unvisited := slices.DeleteFunc(neighbours, w.IsVisited)
		candidates := slice.Map(unvisited, func(n T) candidate[T] {
			return candidate[T]{
				coord:  n,
				cost:   c.cost + 1,
				parent: &c,
			}
		})

		w.candidates.Push(candidates...)
	}

	return []T{}
}

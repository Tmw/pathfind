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
	eventlog   []Event

	MaxCost int
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

func (w *BFS[T]) publish(e Event) {
	w.eventlog = append(w.eventlog, e)
}

func (w *BFS[T]) EventLog() []Event {
	return w.eventlog
}

func (w *BFS[T]) Walk() []T {
	for w.candidates.Len() > 0 {
		c := w.candidates.Pop()
		w.publish(EventCandidateVisited[T]{CandidateID: c.coord})

		if w.MaxCost > 0 && c.cost >= w.MaxCost {
			w.publish(EventMaxCostReached{})
			break
		}

		if w.adapter.IsFinish(c.coord) {
			path := backtrace[T](c)
			w.publish(EventFinishReached[T]{Path: path})
			return path
		}

		w.visit(c.coord)

		neighbours := w.adapter.Neighbours(c.coord)
		unvisited := slices.DeleteFunc(neighbours, w.IsVisited)

		for idx := range unvisited {
			w.publish(EventCandidateAdded[T]{CandidateID: unvisited[idx]})
		}

		candidates := slice.Map(unvisited, func(n T) candidate[T] {
			return candidate[T]{
				coord:  n,
				cost:   c.cost + 1,
				parent: &c,
			}
		})

		w.candidates.Push(candidates...)
	}

	w.publish(EventUnsolvable{})
	return []T{}
}

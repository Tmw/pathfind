package pathfind

import (
	"github.com/tmw/pathfind/pkg/prioqueue"
)

type AStar[T comparable] struct {
	candidates *prioqueue.Prioqueue[candidate[T]]
	visited    map[T]struct{}
	adapter    Adapter[T]
	eventlog   []Event

	MaxCost int
}

func NewAStar[T comparable](start T, adapter Adapter[T]) AStar[T] {
	sc := candidate[T]{
		coord:  start,
		parent: nil,
	}

	candidates := prioqueue.New[candidate[T]]()
	candidates.Push(sc, adapter.CostToFinish(start))

	return AStar[T]{
		candidates: candidates,
		visited:    make(map[T]struct{}),
		adapter:    adapter,
	}
}

func (w *AStar[T]) isVisited(c T) bool {
	_, v := w.visited[c]
	return v
}

func (w *AStar[T]) visit(c T) {
	w.visited[c] = struct{}{}
}

func (w *AStar[T]) publish(e Event) {
	w.eventlog = append(w.eventlog, e)
}

func (w *AStar[T]) EventLog() []Event {
	return w.eventlog
}

func (w *AStar[T]) Walk() []T {
	for w.candidates.Len() > 0 {
		currentNode := w.candidates.Pop()

		if w.MaxCost > 0 && currentNode.cost >= w.MaxCost {
			w.publish(EventMaxCostReached{})
			break
		}

		if w.adapter.IsFinish(currentNode.coord) {
			path := backtrace[T](currentNode)
			w.publish(EventFinishReached[T]{Path: path})
			return path
		}

		neighbours := w.adapter.Neighbours(currentNode.coord)
		for _, n := range neighbours {
			if w.isVisited(n) {
				continue
			}

			newCandidate := candidate[T]{
				coord:  n,
				parent: &currentNode,
				cost:   currentNode.cost + 1,
			}

			predicate := func(i candidate[T]) bool {
				return i.coord == newCandidate.coord
			}

			existingCandidateIdx := w.candidates.IndexFunc(predicate)
			if existingCandidateIdx > 0 {
				existingCandidate := w.candidates.PeekItem(existingCandidateIdx)
				newCost := w.adapter.CostToFinish(n) + existingCandidate.cost
				if newCost < w.candidates.PriorityOfItem(existingCandidateIdx) {
					w.candidates.UpdateAtIndex(existingCandidateIdx, newCandidate, newCost)
				}
			} else {
				neighbourCost := w.adapter.CostToFinish(n) + newCandidate.cost
				w.candidates.Push(newCandidate, neighbourCost)
				w.publish(EventCandidateAdded[T]{CandidateID: newCandidate.coord})
			}
		}

		w.visit(currentNode.coord)
		w.publish(EventCandidateVisited[T]{CandidateID: currentNode.coord})
	}

	w.publish(EventUnsolvable{})
	return []T{}
}

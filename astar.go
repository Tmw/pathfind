package pathfind

import (
	"github.com/tmw/pathfind/pkg/prioqueue"
)

type candidate[T comparable] struct {
	coord  T
	parent *candidate[T]
	cost   int
}

type AStar[T comparable] struct {
	candidates *prioqueue.Prioqueue[candidate[T]]
	visited    map[T]struct{}
	adapter    Adapter[T]
}

func NewAStar[T comparable](start T, adapter Adapter[T]) AStar[T] {
	sc := candidate[T]{
		coord:  start,
		parent: nil,
	}

	candidates := prioqueue.New[candidate[T]]()
	candidates.Push(sc, adapter.DistanceToFinish(start))

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

func (w *AStar[T]) isWalkable(c T) bool {
	return w.adapter.IsCellWalkable(c)
}

func (w *AStar[T]) isFinish(c T) bool {
	return w.adapter.IsCellFinish(c)
}

func (w *AStar[T]) Walk() []T {
	path := []T{}

	for w.candidates.Len() > 0 {
		currentNode := w.candidates.Pop()

		if w.isFinish(currentNode.coord) {
			hop := currentNode
			for hop.parent != nil {
				path = append(path, hop.coord)
				hop = *hop.parent
			}

			return path
		}

		neighbours := w.adapter.Neighbours(currentNode.coord)
		for _, n := range neighbours {
			if w.isVisited(n) || !w.isWalkable(n) {
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
				newCost := w.adapter.DistanceToFinish(n) + existingCandidate.cost
				if newCost < w.candidates.PriorityOfItem(existingCandidateIdx) {
					w.candidates.UpdateAtIndex(existingCandidateIdx, newCandidate, newCost)
				}
			} else {
				neighbourCost := w.adapter.DistanceToFinish(n) + newCandidate.cost
				w.candidates.Push(newCandidate, neighbourCost)
			}
		}
		w.visited[currentNode.coord] = struct{}{}
	}

	return path
}

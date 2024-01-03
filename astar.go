package pathfind

import (
	"github.com/tmw/pathfind/pkg/prioqueue"
)

type astar[T comparable] struct {
	candidates *prioqueue.Prioqueue[candidate[T]]
	start      T
}

func newAStar[T comparable](start T) *astar[T] {
	return &astar[T]{
		start:      start,
		candidates: prioqueue.New[candidate[T]](),
	}
}

func (w *astar[T]) Walk(ctx SolveContext[T]) []T {
	// add initial starting position
	sc := candidate[T]{
		coord:  w.start,
		parent: nil,
	}
	w.candidates.Push(sc, ctx.Adapter().CostToFinish(w.start))

	// main loop
	for w.candidates.Len() > 0 {
		currentNode := w.candidates.Pop()

		if ctx.MaxCost > 0 && currentNode.cost >= ctx.MaxCost {
			ctx.Publish(EventMaxCostReached{})
			break
		}

		if ctx.Adapter().IsFinish(currentNode.coord) {
			path := backtrace[T](currentNode)
			ctx.Publish(EventFinishReached[T]{Path: path})
			return path
		}

		neighbours := ctx.Adapter().Neighbours(currentNode.coord)
		for _, n := range neighbours {
			if ctx.IsVisited(n) {
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
				newCost := ctx.Adapter().CostToFinish(n) + existingCandidate.cost
				if newCost < w.candidates.PriorityOfItem(existingCandidateIdx) {
					w.candidates.UpdateAtIndex(existingCandidateIdx, newCandidate, newCost)
				}
			} else {
				neighbourCost := ctx.Adapter().CostToFinish(n) + newCandidate.cost
				w.candidates.Push(newCandidate, neighbourCost)
				ctx.Publish(EventCandidateAdded[T]{CandidateID: newCandidate.coord})
			}
		}

		ctx.Visit(currentNode.coord)
		ctx.Publish(EventCandidateVisited[T]{CandidateID: currentNode.coord})
	}

	ctx.Publish(EventUnsolvable{})
	return []T{}
}

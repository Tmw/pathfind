package pathfind

import (
	"slices"

	"github.com/tmw/pathfind/pkg/queue"
	"github.com/tmw/pathfind/pkg/slice"
)

type bfs[T comparable] struct {
	candidates queue.Queue[candidate[T]]
}

func newBFS[T comparable](start T) *bfs[T] {
	sc := candidate[T]{
		coord:  start,
		parent: nil,
	}

	candidates := queue.New[candidate[T]](sc)

	return &bfs[T]{
		candidates: candidates,
	}
}

func (w *bfs[T]) Walk(ctx SolveContext[T]) []T {
	for w.candidates.Len() > 0 {
		c := w.candidates.Pop()

		// in case we enqueued the same node multiple times.
		if ctx.IsVisited(c.coord) {
			continue
		}

		ctx.Publish(EventCandidateVisited[T]{CandidateID: c.coord})

		if ctx.MaxCost > 0 && c.cost >= ctx.MaxCost {
			ctx.Publish(EventMaxCostReached{})
			break
		}

		if ctx.Adapter().IsFinish(c.coord) {
			path := backtrace[T](c)
			ctx.Publish(EventFinishReached[T]{Path: path})
			return path
		}

		ctx.Visit(c.coord)

		neighbours := ctx.Adapter().Neighbours(c.coord)
		unvisited := slices.DeleteFunc(neighbours, ctx.IsVisited)

		for idx := range unvisited {
			ctx.Publish(EventCandidateAdded[T]{CandidateID: unvisited[idx]})
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

	ctx.Publish(EventUnsolvable{})
	return []T{}
}

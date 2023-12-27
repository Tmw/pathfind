package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/tmw/pathfind/pkg/arena"
	"github.com/tmw/pathfind/pkg/prioqueue"
)

const input = `
##############################
#............................#
#..S.........................#
#######.#############........#
#............................#
#......#####################.#
#............................#
########.###############.....#
#............................#
#......########..............#
#......#......################
#..........#.................#
############.....#...........#
#......#...#.....#...........#
#......#...#.....#......F....#
#......#...#.....#...........#
##############################
`

type candidate[T comparable] struct {
	coord  T
	parent *candidate[T]
	cost   int
}

type Adapter[T comparable] interface {
	Neighbours(T) []T
	DistanceToFinish(T) int
	IsCellWalkable(T) bool
	IsCellFinish(T) bool
}

type FuncAdapter[T comparable] struct {
	NeighboursFn       func(T) []T
	DistanceToFinishFn func(T) int
	IsCellWalkableFn   func(T) bool
	IsCellFinishFn     func(T) bool
}

func (a *FuncAdapter[T]) Neighbours(c T) []T {
	return a.NeighboursFn(c)
}

func (a *FuncAdapter[T]) DistanceToFinish(c T) int {
	return a.DistanceToFinishFn(c)
}

func (a *FuncAdapter[T]) IsCellWalkable(c T) bool {
	return a.IsCellWalkableFn(c)
}

func (a *FuncAdapter[T]) IsCellFinish(c T) bool {
	return a.IsCellFinishFn(c)
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

func main() {
	m, err := arena.Parse(input)
	if err != nil {
		log.Fatal(err)
	}

	w := NewAStar[arena.Coordinate](m.StartCoordinate(), &FuncAdapter[arena.Coordinate]{
		NeighboursFn: func(c arena.Coordinate) []arena.Coordinate {
			return m.NeighboursOfCoordinate(c)
		},

		DistanceToFinishFn: func(c arena.Coordinate) int {
			return c.DistanceTo(m.FinishCoordinate())
		},

		IsCellWalkableFn: func(c arena.Coordinate) bool {
			return m.CellTypeForCoordinate(c) != arena.CellTypeNonWalkable
		},

		IsCellFinishFn: func(c arena.Coordinate) bool {
			return m.CellTypeForCoordinate(c) == arena.CellTypeFinish
		},
	})

	start := time.Now()
	path := w.Walk()
	d := time.Since(start)

	if path != nil {
		fmt.Printf("\033[H\033[2J")
		m.RenderWithPath(os.Stdout, path)
	}

	fmt.Printf("\n\nsolve duration: %v\n", d)
	fmt.Printf("visited count: %v\n", len(w.visited))
}

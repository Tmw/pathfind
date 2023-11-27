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

type candidate struct {
	coord  arena.Coordinate
	parent *candidate
	cost   int
}

type walker struct {
	candidates *prioqueue.Prioqueue[candidate]
	visited    map[arena.Coordinate]struct{}

	// TODO: Instead of taking these as functions, perhaps make this an interface?
	neighboursFn            func(arena.Coordinate) []arena.Coordinate
	distanceToFinishFn      func(arena.Coordinate) int
	cellTypeForCoordinateFn func(arena.Coordinate) arena.CellType
}

type path []arena.Coordinate

func (w *walker) isVisited(c arena.Coordinate) bool {
	_, v := w.visited[c]
	return v
}

func (w *walker) isWalkable(c arena.Coordinate) bool {
	return w.cellTypeForCoordinateFn(c) == arena.CellTypeNonWalkable
}

func (w *walker) isFinish(c arena.Coordinate) bool {
	return w.cellTypeForCoordinateFn(c) == arena.CellTypeFinish
}

func (w *walker) Walk() path {
	path := []arena.Coordinate{}

	for w.candidates.Len() > 0 {
		currentNode := w.candidates.PopValue()

		if w.isFinish(currentNode.coord) {
			hop := currentNode
			for hop.parent != nil {
				path = append(path, hop.coord)
				hop = *hop.parent
			}

			return path
		}

		neighbours := w.neighboursFn(currentNode.coord)
		for _, n := range neighbours {
			if w.isVisited(n) || !w.isWalkable(n) {
				continue
			}

			nn := candidate{
				coord:  n,
				parent: &currentNode,
				cost:   currentNode.cost + 1,
			}

			predicate := func(i candidate) bool {
				return i.coord == nn.coord
			}

			existingCandidateIdx := w.candidates.IndexFunc(predicate)
			if existingCandidateIdx > 0 {
				existingCandidate := w.candidates.PeekItem(existingCandidateIdx)
				if existingCandidate != nil {
					newCost := w.distanceToFinishFn(n) + existingCandidate.Value.cost
					if newCost < existingCandidate.Priority() {
						w.candidates.UpdateAtIndex(existingCandidateIdx, nn, newCost)
					}
				}
			} else {
				// TODO: Abstract into its own cost function so we can reuse.
				neighbourCost := w.distanceToFinishFn(n) + nn.cost
				w.candidates.Push(nn, neighbourCost)
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

	n := candidate{
		coord:  m.StartCoordinate(),
		parent: nil,
	}

	c := prioqueue.New[candidate]()
	c.Push(n, m.StartCoordinate().DistanceTo(m.FinishCoordinate()))

	w := walker{
		candidates: c,
		visited:    map[arena.Coordinate]struct{}{},

		// functions to interface with the arena
		neighboursFn: func(c arena.Coordinate) []arena.Coordinate {
			return m.NeighboursOfCoordinate(c)
		},

		distanceToFinishFn: func(c arena.Coordinate) int {
			return c.DistanceTo(m.FinishCoordinate())
		},

		cellTypeForCoordinateFn: func(c arena.Coordinate) arena.CellType {
			return m.CellTypeForCoordinate(c)
		},
	}

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

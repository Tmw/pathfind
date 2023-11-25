package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"time"

	"github.com/tmw/pathfind/pkg/arena"
	"github.com/tmw/pathfind/pkg/prioqueue"
)

const input = `
##############################
#............................#
#..S.........................#
#####################........#
#............................#
#......#######################
#............................#
########################.....#
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

type node struct {
	coord  arena.Coordinate
	parent *node
}

type walker struct {
	candidates *prioqueue.Prioqueue[node]
	visited    []arena.Coordinate
	path       *[]arena.Coordinate

	// TODO: Instead of taking these as functions, perhaps make this an interface?
	neighboursFn            func(arena.Coordinate) []arena.Coordinate
	distanceToFinishFn      func(arena.Coordinate) int
	cellTypeForCoordinateFn func(arena.Coordinate) *arena.CellType
}

type done bool

var visitedLen int

func (w *walker) step() done {
	var (
		nonWalkable = arena.CellTypeNonWalkable
		finish      = arena.CellTypeFinish
	)

	if w.candidates.Len() <= 0 {
		return true
	}

	item := w.candidates.PopItem()
	currentNode, _ := item.Value, item.Priority()

	if t := w.cellTypeForCoordinateFn(currentNode.coord); t != nil && *t == finish {

		w.path = &[]arena.Coordinate{}

		var hop node = currentNode
		for hop.parent != nil {
			*w.path = append(*w.path, hop.coord)
			hop = *hop.parent
		}

		// we found the finish, now backtrack.
		return true
	}

	neighbours := w.neighboursFn(currentNode.coord)
	for _, n := range neighbours {
		isAlreadyVisited := slices.Contains(w.visited, n)
		isNonWalkable := *w.cellTypeForCoordinateFn(n) == nonWalkable

		if isAlreadyVisited || isNonWalkable {
			continue
		}

		// TODO: I think in order to make it actually find the quickest path,
		// it'll need to update the candidate if it's already in the heap,
		// and update its cost and parent node.
		nn := node{
			coord:  n,
			parent: &currentNode,
		}

		predicate := func(i node) bool {
			return i.coord == nn.coord
		}

		if !w.candidates.ContainsFunc(nn, predicate) {
			// push candidate to the list.
			neighbourCost := w.distanceToFinishFn(n)
			// if candidate exist, we should update its cost.
			w.candidates.Push(nn, neighbourCost)
		}
	}

	w.visited = append(w.visited, currentNode.coord)
	visitedLen = len(w.visited)

	return false
}

func main() {
	m, err := arena.Parse(input)
	if err != nil {
		log.Fatal(err)
	}

	n := node{
		coord:  m.StartCoordinate(),
		parent: nil,
	}

	c := prioqueue.New[node]()
	c.Push(n, m.StartCoordinate().DistanceTo(m.FinishCoordinate()))

	w := walker{
		candidates: c,
		visited:    []arena.Coordinate{},

		// functions to interface with the arena
		neighboursFn: func(c arena.Coordinate) []arena.Coordinate {
			return m.NeighboursOfCoordinate(c)
		},

		distanceToFinishFn: func(c arena.Coordinate) int {
			return c.DistanceTo(m.FinishCoordinate())
		},

		cellTypeForCoordinateFn: func(c arena.Coordinate) *arena.CellType {
			return m.CellTypeForCoordinate(c)
		},
	}

	for !w.step() {
		fmt.Printf("\033[H\033[2J")
		m.RenderWithVisited(os.Stdout, w.visited)
		fmt.Printf("\n\n-> visitedLen=%+v\n", visitedLen)
		fmt.Printf("-> candidatesLen=%+v\n", w.candidates.Len())

		time.Sleep(40 * time.Millisecond)
	}

	if w.path != nil {
		fmt.Printf("\033[H\033[2J")
		m.RenderWithPath(os.Stdout, *w.path)
	}
}

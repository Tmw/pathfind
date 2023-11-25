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

type walker struct {
	candidates *prioqueue.Prioqueue[arena.Coordinate]
	visited    []arena.Coordinate

	// TODO: Instead of taking these as functions, perhaps make this an interface?
	neighboursFn            func(arena.Coordinate) []arena.Coordinate
	distanceToFinishFn      func(arena.Coordinate) int
	cellTypeForCoordinateFn func(arena.Coordinate) *arena.CellType
}

type done bool

var looking arena.Coordinate
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
	coord, _ := item.Value, item.Priority()
	looking = item.Value

	if t := w.cellTypeForCoordinateFn(coord); t != nil && *t == finish {
		return true
	}

	neighbours := w.neighboursFn(coord)
	for _, n := range neighbours {
		isAlreadyVisited := slices.Contains(w.visited, n)
		isNonWalkable := *w.cellTypeForCoordinateFn(n) == nonWalkable

		if isAlreadyVisited || isNonWalkable {
			continue
		}

		// TODO: I think in order to make it actually find the quickest path,
		// it'll need to update the candidate if it's already in the heap,
		// and update its cost.
		// and in order to be able to return and mark a path, we'll need
		// keep track of where it came from. So for each candidate, I think we'll need to store its "parent"
		// as well?
		if !w.candidates.Contains(n) {
			// push candidate to the list.
			neighbourCost := w.distanceToFinishFn(n)
			// if candidate exist, we should update its cost.
			w.candidates.Push(n, neighbourCost)
		}
	}

	w.visited = append(w.visited, coord)
	visitedLen = len(w.visited)

	return false
}

func costForCoordinate(c, start, finish arena.Coordinate) int {
	return c.DistanceTo(finish)
}

func main() {
	m, err := arena.Parse(input)
	if err != nil {
		log.Fatal(err)
	}

	c := prioqueue.New[arena.Coordinate]()
	c.Push(m.StartCoordinate(), costForCoordinate(m.StartCoordinate(), m.StartCoordinate(), m.FinishCoordinate()))

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
		m.RenderWithVisited(os.Stdout, w.visited)
		fmt.Printf("\n\nlooking=%+v\n", looking)
		fmt.Printf("visitedLen=%+v\n", visitedLen)
		fmt.Printf("candidatesLen=%+v\n", w.candidates.Len())
		fmt.Printf("distance=%+v\n", looking.DistanceTo(m.FinishCoordinate()))

		time.Sleep(80 * time.Millisecond)
		fmt.Printf("\033[H\033[2J")
		// clear the screen?
		// render out the thing?
		// sleep for a bit?
	}

	// pq := prioqueue.New[string]()
	//
	// pq.Push("wut 10", 10)
	// pq.Push("wut 11", 11)
	// pq.Push("wut 8", 8)
	//
	// for pq.Len() > 0 {
	// 	item := pq.PopItem()
	// 	fmt.Printf("%s => %d\n", item.Value, item.Priority())
	// }
	//
	// fmt.Print(Floor)
}

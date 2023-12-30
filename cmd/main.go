package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/tmw/pathfind"
	"github.com/tmw/pathfind/pkg/arena"
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

func main() {
	m, err := arena.Parse(input)
	if err != nil {
		log.Fatal(err)
	}

	w := pathfind.NewAStar[arena.Coordinate](m.StartCoordinate(), &pathfind.FuncAdapter[arena.Coordinate]{
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
}

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"time"

	"github.com/tmw/pathfind"
	"github.com/tmw/pathfind/pkg/arena"
)

var (
	filename  string
	algorithm string
	verbose   bool

	// configure map symbols
	symbolNonWalkable string
	symbolWalkable    string
	symbolStart       string
	symbolFinish      string
	symbolPath        string
)

func init() {
	flag.StringVar(&filename, "filename", "", "path of the file to read")
	flag.StringVar(&symbolNonWalkable, "symbolNonWalkable", "", "symbol for tile of type non-walkable")
	flag.StringVar(&symbolWalkable, "symbolWalkable", "", "symbol for tile of type walkable")
	flag.StringVar(&symbolStart, "symbolStart", "", "symbol for tile of type start")
	flag.StringVar(&symbolFinish, "symbolFinish", "", "symbol for tile of type finish")
	flag.StringVar(&symbolPath, "symbolPath", "", "symbol for tile of type path")
	flag.StringVar(&algorithm, "algorithm", "astar", "algorithm to use. either astar or bfs are supported")
	flag.BoolVar(&verbose, "verbose", true, "print runtime information")
}

func main() {
	flag.Parse()

	if !slices.Contains([]string{"astar", "bfs"}, algorithm) {
		log.Fatal("provided algorithm not supported, must be any of: astar, bfs")
	}

	contents, err := getContents()
	if err != nil {
		log.Fatal(err)
	}

	assignSymbols()

	if err := solve(contents); err != nil {
		log.Fatal(err)
	}
}

func getAlgorithm() pathfind.Algorithm {
	if algorithm == "bfs" {
		return pathfind.AlgorithmBFS
	}

	return pathfind.AlgorithmAStar
}

func assignSymbols() {
	if len(symbolNonWalkable) > 0 {
		arena.SymbolNonWalkable = symbolNonWalkable
	}

	if len(symbolWalkable) > 0 {
		arena.SymbolWalkable = symbolWalkable
	}

	if len(symbolStart) > 0 {
		arena.SymbolStart = symbolStart
	}

	if len(symbolFinish) > 0 {
		arena.SymbolFinish = symbolFinish
	}

	if len(symbolPath) > 0 {
		arena.SymbolPath = symbolPath
	}
}

func getContents() (string, error) {
	if len(filename) > 0 {
		return readfile(filename)
	}

	return readStdin()
}

func readStdin() (string, error) {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func readfile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func solve(input string) error {
	a, err := arena.Parse(input)
	if err != nil {
		return err
	}

	s := pathfind.NewSolver[arena.Coordinate](
		getAlgorithm(),
		a.StartCoordinate(),
		&pathfind.FuncAdapter[arena.Coordinate]{
			NeighboursFn: func(c arena.Coordinate) []arena.Coordinate {
				n := a.NeighboursOfCoordinate(c)
				return slices.DeleteFunc(n, func(c arena.Coordinate) bool {
					return a.CellTypeForCoordinate(c) == arena.CellTypeNonWalkable
				})
			},

			CostToFinishFn: func(c arena.Coordinate) int {
				return c.DistanceTo(a.FinishCoordinate())
			},

			IsFinishFn: func(c arena.Coordinate) bool {
				return a.CellTypeForCoordinate(c) == arena.CellTypeFinish
			},
		},
	)

	s.MaxCost = 50
	start := time.Now()
	path := s.Walk()
	duration := time.Since(start)

	if path != nil {
		a.RenderWithPath(os.Stdout, path)
		fmt.Print("\n\n")
	}

	if verbose {
		candidates := make(map[arena.Coordinate]struct{})
		for _, e := range s.EventLog() {
			if cv, ok := e.(pathfind.EventCandidateVisited[arena.Coordinate]); ok {
				candidates[cv.CandidateID] = struct{}{}
			}
		}

		fmt.Printf("used algorithm: \t\t%s\n", algorithm)
		fmt.Printf("duration: \t\t\t%s\n", duration)
		fmt.Printf("unique candidates visited: \t%d\n", len(candidates))
	}

	return nil
}

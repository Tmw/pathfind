package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/tmw/pathfind"
	"github.com/tmw/pathfind/pkg/arena"
)

var (
	filename string

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
}

func main() {
	flag.Parse()

	contents, err := getContents()
	if err != nil {
		log.Fatal(err)
	}

	assignSymbols()

	if err := solve(contents); err != nil {
		log.Fatal(err)
	}
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

	w := pathfind.NewAStar[arena.Coordinate](a.StartCoordinate(), &pathfind.FuncAdapter[arena.Coordinate]{
		NeighboursFn: func(c arena.Coordinate) []arena.Coordinate {
			return a.NeighboursOfCoordinate(c)
		},

		DistanceToFinishFn: func(c arena.Coordinate) int {
			return c.DistanceTo(a.FinishCoordinate())
		},

		IsCellWalkableFn: func(c arena.Coordinate) bool {
			return a.CellTypeForCoordinate(c) != arena.CellTypeNonWalkable
		},

		IsCellFinishFn: func(c arena.Coordinate) bool {
			return a.CellTypeForCoordinate(c) == arena.CellTypeFinish
		},
	})

	path := w.Walk()

	if path != nil {
		a.RenderWithPath(os.Stdout, path)
	}

	return nil
}

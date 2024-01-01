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
)

func init() {
	flag.StringVar(&filename, "filename", "", "path of the file to read")
}

func main() {
	flag.Parse()

	contents, err := getContents()
	if err != nil {
		log.Fatal(err)
	}

	if err := solve(contents); err != nil {
		log.Fatal(err)
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
	m, err := arena.Parse(input)
	if err != nil {
		return err
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

	path := w.Walk()

	if path != nil {
		m.RenderWithPath(os.Stdout, path)
	}

	return nil
}

package arena

import (
	"errors"
	"strings"

	"github.com/tmw/pathfind/pkg/slice"
)

var (
	ErrorInvalidArenaNoStart        = errors.New("invalid map: no start cell found")
	ErrorInvalidArenaNoFinish       = errors.New("invalid map: no finish cell found")
	ErrorInvalidArenaMultipleStart  = errors.New("invalid map: multiple start cells found")
	ErrorInvalidArenaMultipleFinish = errors.New("invalid map: multiple finish cells found")
)

var (
	SymbolNonWalkable = "#"
	SymbolWalkable    = "."
	SymbolStart       = "S"
	SymbolFinish      = "F"
	SymbolPath        = "@"
)

func Parse(input string) (*Arena, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	cells := make([][]CellType, len(lines))
	for i := range cells {
		cells[i] = slice.Map(strings.Split(lines[i], ""), symbolToType)
	}

	start, stop, err := findStartAndFinish(cells)
	if err != nil {
		return nil, err
	}

	m := &Arena{
		cells:      cells,
		startCell:  *start,
		finishCell: *stop,
	}

	return m, nil
}

func findStartAndFinish(cells [][]CellType) (*Coordinate, *Coordinate, error) {
	var (
		start  *Coordinate
		finish *Coordinate
	)

	for y := range cells {
		for x := range cells[y] {
			cell := cells[y][x]
			if cell == CellTypeStart {
				if start != nil {
					return start, finish, ErrorInvalidArenaMultipleStart
				}

				start = &Coordinate{x: x, y: y}
			}

			if cell == CellTypeFinish {
				if finish != nil {
					return start, finish, ErrorInvalidArenaMultipleFinish
				}

				finish = &Coordinate{x: x, y: y}
			}
		}
	}

	if start == nil {
		return start, finish, ErrorInvalidArenaNoStart
	}

	if finish == nil {
		return start, finish, ErrorInvalidArenaNoFinish
	}

	return start, finish, nil
}

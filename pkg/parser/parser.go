package parser

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/tmw/pathfind/pkg/slice"
)

var (
	InvalidMapNoStart        = errors.New("invalid map: no start cell found")
	InvalidMapNoFinish       = errors.New("invalid map: no finish cell found")
	InvalidMapMultipleStart  = errors.New("invalid map: multiple start cells found")
	InvalidMapMultipleFinish = errors.New("invalid map: multiple finish cells found")
)

const (
	SymbolNonWalkable = "#"
	SymbolWalkable    = "."
	SymbolStart       = "S"
	SymbolFinish      = "F"
	SymbolPath        = "@"
)

type CellType uint8

func (t CellType) String() string {
	switch t {
	case CellTypeNonWalkable:
		return SymbolNonWalkable

	case CellTypeStart:
		return SymbolStart

	case CellTypeFinish:
		return SymbolFinish

	case CellTypeWalkable:
		return SymbolWalkable

	default:
		return SymbolWalkable
	}
}

const (
	CellTypeWalkable CellType = iota
	CellTypeNonWalkable
	CellTypeStart
	CellTypeFinish
	TyileTypePath
)

type Coordinate struct {
	x, y int
}

type Map struct {
	cells     [][]CellType
	startCell Coordinate
	endCell   Coordinate
}

// render the map into the writer
func (m Map) Render(w io.Writer) {
	for x := range m.cells {
		if x > 0 {
			fmt.Fprintf(w, "\n")
		}

		for y := range m.cells[x] {
			fmt.Fprintf(w, m.cells[x][y].String())
		}
	}
}

func Parse(input string) (*Map, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	cells := make([][]CellType, len(lines))
	for i := range cells {
		cells[i] = slice.Map(strings.Split(lines[i], ""), symbolToType)
	}

	start, stop, err := findStartAndFinish(cells)
	if err != nil {
		return nil, err
	}

	m := &Map{
		cells:     cells,
		startCell: *start,
		endCell:   *stop,
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
					return start, finish, InvalidMapMultipleStart
				}

				start = &Coordinate{x: x, y: y}
			}

			if cell == CellTypeFinish {
				if finish != nil {
					return start, finish, InvalidMapMultipleFinish
				}

				finish = &Coordinate{x: x, y: y}
			}
		}
	}

	if start == nil {
		return start, finish, InvalidMapNoStart
	}

	if finish == nil {
		return start, finish, InvalidMapNoFinish
	}

	return start, finish, nil
}

func symbolToType(i string) CellType {
	switch i {
	case SymbolNonWalkable:
		return CellTypeNonWalkable

	case SymbolStart:
		return CellTypeStart

	case SymbolFinish:
		return CellTypeFinish

	case SymbolWalkable:
		return CellTypeWalkable

	default:
		return CellTypeWalkable
	}
}

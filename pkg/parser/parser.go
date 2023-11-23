package parser

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/tmw/pathfind/pkg/slice"
)

var (
	InvalidMapNoStart       = errors.New("invalid map: no start tile found")
	InvalidMapNoEnd         = errors.New("invalid map: no end tile found")
	InvalidMapMultipleStart = errors.New("invalid map: multiple start tiles found")
	InvalidMapMultipleEnd   = errors.New("invalid map: multiple end tiles found")
)

const (
	SymbolNonWalkable = "#"
	SymbolWalkable    = "."
	SymbolStart       = "S"
	SymbolEnd         = "E"
	SymbolPath        = "@"
)

type TileType uint8

func (t TileType) String() string {
	switch t {
	case TileTypeNonWalkable:
		return SymbolNonWalkable

	case TileTypeStart:
		return SymbolStart

	case TileTypeEnd:
		return SymbolEnd

	case TileTypeWalkable:
		return SymbolWalkable

	default:
		return SymbolWalkable
	}
}

const (
	TileTypeWalkable TileType = iota
	TileTypeNonWalkable
	TileTypeStart
	TileTypeEnd
	TyileTypePath
)

type Map [][]TileType

// render the map into the writer
func (m Map) Render(w io.Writer) {
	for x := range m {
		if x > 0 {
			fmt.Fprintf(w, "\n")
		}

		for y := range m[x] {
			fmt.Fprintf(w, m[x][y].String())
		}
	}
}

func Parse(floor string) (Map, error) {
	if err := valdiateInput(floor); err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(floor), "\n")
	m := make([][]TileType, len(lines))
	for i := range m {
		m[i] = slice.Map(strings.Split(lines[i], ""), symbolToType)
	}

	return m, nil
}

func valdiateInput(input string) error {
	var (
		startCount = strings.Count(input, SymbolStart)
		endCount   = strings.Count(input, SymbolEnd)
	)

	if startCount == 0 {
		return InvalidMapNoStart
	}

	if startCount > 1 {
		return InvalidMapMultipleStart
	}

	if endCount == 0 {
		return InvalidMapNoEnd
	}

	if endCount > 1 {
		return InvalidMapMultipleEnd
	}

	return nil
}

func symbolToType(i string) TileType {
	switch i {
	case SymbolNonWalkable:
		return TileTypeNonWalkable

	case SymbolStart:
		return TileTypeStart

	case SymbolEnd:
		return TileTypeEnd

	case SymbolWalkable:
		return TileTypeWalkable

	default:
		return TileTypeWalkable
	}
}

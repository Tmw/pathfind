package arena

import (
	"fmt"
	"io"
	"slices"
)

type Arena struct {
	cells      [][]CellType
	startCell  Coordinate
	finishCell Coordinate
}

// render the map into the writer
func (m *Arena) Render(w io.Writer) {
	for x := range m.cells {
		if x > 0 {
			fmt.Fprintf(w, "\n")
		}

		for y := range m.cells[x] {
			fmt.Fprintf(w, "%s", m.cells[x][y])
		}
	}
}

func (m *Arena) RenderWithVisited(w io.Writer, visited []Coordinate) {
	for y := range m.cells {
		if y > 0 {
			fmt.Fprintf(w, "\n")
		}

		for x := range m.cells[y] {
			c := Coordinate{x: x, y: y}
			if slices.Contains(visited, c) {
				fmt.Fprint(w, "v")
			} else {
				fmt.Fprintf(w, "%s", m.cells[y][x])
			}
		}
	}
}

func (m *Arena) RenderWithPath(w io.Writer, path []Coordinate) {
	for y := range m.cells {
		if y > 0 {
			fmt.Fprintf(w, "\n")
		}

		for x := range m.cells[y] {
			c := Coordinate{x: x, y: y}
			if slices.Contains(path, c) {
				fmt.Fprintf(w, "%s", SymbolPath)
			} else {
				fmt.Fprintf(w, "%s", m.cells[y][x])
			}
		}
	}
}

func (m *Arena) NeighboursOfCoordinate(c Coordinate) []Coordinate {
	neighbours := []Coordinate{}

	if n := c.North(); m.CellTypeForCoordinate(n) != CellTypeUndefined {
		neighbours = append(neighbours, n)
	}

	if n := c.West(); m.CellTypeForCoordinate(n) != CellTypeUndefined {
		neighbours = append(neighbours, n)
	}

	if n := c.South(); m.CellTypeForCoordinate(n) != CellTypeUndefined {
		neighbours = append(neighbours, n)
	}

	if n := c.East(); m.CellTypeForCoordinate(n) != CellTypeUndefined {
		neighbours = append(neighbours, n)
	}

	return neighbours
}

func (m *Arena) CellTypeForCoordinate(c Coordinate) CellType {
	if c.y < 0 {
		return CellTypeUndefined
	}
	if c.y >= len(m.cells) {
		return CellTypeUndefined
	}
	if c.x < 0 {
		return CellTypeUndefined
	}
	if c.x >= len(m.cells[c.y]) {
		return CellTypeUndefined
	}

	return m.cells[c.y][c.x]
}

func (m *Arena) StartCoordinate() Coordinate {
	return m.startCell
}

func (m *Arena) FinishCoordinate() Coordinate {
	return m.finishCell
}

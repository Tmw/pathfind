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
			fmt.Fprintf(w, m.cells[x][y].String())
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
				fmt.Fprintf(w, m.cells[y][x].String())
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
				fmt.Fprint(w, "@")
			} else {
				fmt.Fprintf(w, m.cells[y][x].String())
			}
		}
	}
}

func (m *Arena) safeGetCell(c Coordinate) *CellType {
	if c.y < 0 {
		return nil
	}
	if c.y >= len(m.cells) {
		return nil
	}
	if c.x < 0 {
		return nil
	}
	if c.x >= len(m.cells[c.y]) {
		return nil
	}

	return &m.cells[c.y][c.x]
}

func (m *Arena) NeighboursOfCoordinate(c Coordinate) []Coordinate {
	neighbours := []Coordinate{}

	if n := c.North(); m.safeGetCell(n) != nil {
		neighbours = append(neighbours, n)
	}

	if n := c.West(); m.safeGetCell(n) != nil {
		neighbours = append(neighbours, n)
	}

	if n := c.South(); m.safeGetCell(n) != nil {
		neighbours = append(neighbours, n)
	}

	if n := c.East(); m.safeGetCell(n) != nil {
		neighbours = append(neighbours, n)
	}

	return neighbours
}

func (m *Arena) CellTypeForCoordinate(c Coordinate) *CellType {
	return m.safeGetCell(c)
}

func (m *Arena) StartCoordinate() Coordinate {
	return m.startCell
}

func (m *Arena) FinishCoordinate() Coordinate {
	return m.finishCell
}

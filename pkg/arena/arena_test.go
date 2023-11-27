package arena

import (
	"reflect"
	"slices"
	"testing"

	"github.com/tmw/pathfind/pkg/slice"
)

const (
	input = `
#############
#.S.........#
#...........#
#...........#
#.....F.....#
#...........#
#...........#
#############
`
)

func TestNeighboursOfCoordinate(t *testing.T) {
	tests := map[string]struct {
		c Coordinate
		e []Coordinate
	}{
		"in upper left corner": {
			c: NewCoordinate(0, 0),
			e: []Coordinate{
				NewCoordinate(0, 1),
				NewCoordinate(1, 0),
			},
		},
		"in middle of the arena": {
			c: NewCoordinate(4, 4),
			e: []Coordinate{
				NewCoordinate(3, 4),
				NewCoordinate(5, 4),
				NewCoordinate(4, 3),
				NewCoordinate(4, 5),
			},
		},
		"with nevative coordinates": {
			c: NewCoordinate(-5, -5),
			e: []Coordinate{},
		},
		"out of scope coordinates": {
			c: NewCoordinate(50, 200),
			e: []Coordinate{},
		},
		"on upper right corner": {
			c: NewCoordinate(12, 0),
			e: []Coordinate{
				NewCoordinate(11, 0),
				NewCoordinate(12, 1),
			},
		},
		"on bottom right corner": {
			c: NewCoordinate(12, 7),
			e: []Coordinate{
				NewCoordinate(11, 7),
				NewCoordinate(12, 6),
			},
		},
	}

	a, err := Parse(input)
	if err != nil {
		t.Fatal(err)
	}

	for name, td := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := a.NeighboursOfCoordinate(td.c)
			match := slice.All(actual, func(n Coordinate) bool {
				return slices.Contains(td.e, n)
			})

			if !match {
				t.Errorf("expected neighbours of %+v to be: %+v but received: %+v", td.c, td.e, actual)
			}
		})
	}
}

func TestCellTypeForCoordinate(t *testing.T) {
	tests := map[string]struct {
		c Coordinate
		e CellType
	}{
		"cell is non-walkable": {
			c: NewCoordinate(0, 0),
			e: CellTypeNonWalkable,
		},
		"cell is walkable": {
			c: NewCoordinate(1, 3),
			e: CellTypeWalkable,
		},
		"cell is starting point": {
			c: NewCoordinate(2, 1),
			e: CellTypeStart,
		},
		"cell is finish point": {
			c: NewCoordinate(6, 4),
			e: CellTypeFinish,
		},
		"cell does not exist": {
			c: NewCoordinate(-1, -1),
			e: CellTypeUndefined,
		},
		"cell is out of scope": {
			c: NewCoordinate(50, 200),
			e: CellTypeUndefined,
		},
	}

	a, err := Parse(input)
	if err != nil {
		t.Fatal(err)
	}

	for name, td := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := a.CellTypeForCoordinate(td.c)
			if !reflect.DeepEqual(actual, td.e) {
				t.Errorf("expected CellType on coordinate %+v to be %+v but got %+v", td.c, td.e, actual)
			}
		})
	}
}

func TestStartCoordinate(t *testing.T) {
	a, err := Parse(input)
	if err != nil {
		t.Fatal(err)
	}

	expected := NewCoordinate(2, 1)
	actual := a.StartCoordinate()

	if actual != expected {
		t.Errorf("expected StartCoordinate() to return %+v but received %+v", expected, actual)
	}
}

func TestFinishCoordinate(t *testing.T) {
	a, err := Parse(input)
	if err != nil {
		t.Fatal(err)
	}

	expected := NewCoordinate(6, 4)
	actual := a.FinishCoordinate()

	if actual != expected {
		t.Errorf("expected FinishCoordinate() to return %+v but received %+v", expected, actual)
	}
}

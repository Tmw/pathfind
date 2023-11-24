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
	cases := []struct {
		c Coordinate
		e []Coordinate
	}{
		{
			c: NewCoordinate(0, 0),
			e: []Coordinate{
				NewCoordinate(0, 1),
				NewCoordinate(1, 0),
			},
		},
		{
			c: NewCoordinate(4, 4),
			e: []Coordinate{
				NewCoordinate(3, 4),
				NewCoordinate(5, 4),
				NewCoordinate(4, 3),
				NewCoordinate(4, 5),
			},
		},
		{
			c: NewCoordinate(-5, -5),
			e: []Coordinate{},
		},
		{
			c: NewCoordinate(50, 200),
			e: []Coordinate{},
		},
		{
			c: NewCoordinate(12, 0),
			e: []Coordinate{
				NewCoordinate(11, 0),
				NewCoordinate(12, 1),
			},
		},
	}

	a, err := Parse(input)
	if err != nil {
		t.Fatal(err)
	}

	for _, c := range cases {
		actual := a.NeighboursOfCoordinate(c.c)
		match := slice.All(actual, func(n Coordinate) bool {
			return slices.Contains(c.e, n)
		})

		if !match {
			t.Fatalf("expected neighbours of %+v to be: %+v but received: %+v", c.c, c.e, actual)
		}
	}
}

func TestCellTypeForCoordinate(t *testing.T) {
	var (
		nonWalkable = CellTypeNonWalkable
		walkable    = CellTypeWalkable
		start       = CellTypeStart
		finish      = CellTypeFinish
	)

	cases := []struct {
		c Coordinate
		e *CellType
	}{
		{
			c: NewCoordinate(0, 0),
			e: &nonWalkable,
		},
		{
			c: NewCoordinate(1, 3),
			e: &walkable,
		},
		{
			c: NewCoordinate(2, 1),
			e: &start,
		},
		{
			c: NewCoordinate(6, 4),
			e: &finish,
		},
		{
			c: NewCoordinate(-1, -1),
			e: nil,
		},
		{
			c: NewCoordinate(50, 200),
			e: nil,
		},
	}

	a, err := Parse(input)
	if err != nil {
		t.Fatal(err)
	}

	for _, c := range cases {
		actual := a.CellTypeForCoordinate(c.c)
		if !reflect.DeepEqual(actual, c.e) {
			t.Fatalf("expected CellType on coordinate %+v to be %+v but got %+v", c.c, c.e, actual)
		}
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
		t.Fatalf("expected StartCoordinate() to return %+v but received %+v", expected, actual)
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
		t.Fatalf("expected FinishCoordinate() to return %+v but received %+v", expected, actual)
	}
}

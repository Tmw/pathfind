package arena

import (
	"testing"
)

func TestCoordinateDistanceTo(t *testing.T) {
	tests := map[string]struct {
		c1, c2 Coordinate
		d      int
	}{
		"simple neighbour": {
			c1: NewCoordinate(1, 1),
			c2: NewCoordinate(2, 2),
			d:  2,
		},
		"simple neighbour but reversed": {
			c1: NewCoordinate(2, 2),
			c2: NewCoordinate(1, 1),
			d:  2,
		},
		"distance to invalid coordinate": {
			c1: NewCoordinate(5, 5),
			c2: NewCoordinate(-5, -5),
			d:  20,
		},
	}

	for name, td := range tests {
		t.Run(name, func(t *testing.T) {
			actual := td.c1.DistanceTo(td.c2)
			if actual != td.d {
				t.Errorf("expected distance between %+v and %+v to be %d but got %d", td.c1, td.c2, td.d, actual)
			}
		})
	}
}

func TestCoordinateMove(t *testing.T) {
	tests := map[string]struct {
		c  Coordinate
		fn func(Coordinate) Coordinate
		e  Coordinate
	}{
		"north": {
			c:  NewCoordinate(1, 1),
			fn: Coordinate.North,
			e:  NewCoordinate(1, 0),
		},
		"east": {
			c:  NewCoordinate(1, 1),
			fn: Coordinate.East,
			e:  NewCoordinate(2, 1),
		},
		"west": {
			c:  NewCoordinate(1, 1),
			fn: Coordinate.West,
			e:  NewCoordinate(0, 1),
		},
		"south": {
			c:  NewCoordinate(1, 1),
			fn: Coordinate.South,
			e:  NewCoordinate(1, 2),
		},
	}

	for name, td := range tests {
		t.Run(name, func(t *testing.T) {
			actual := td.fn(td.c)
			if actual != td.e {
				t.Fatalf("expected to get %+v when calling %+v on %+v, instead received %+v", td.e, name, td.c, actual)
			}
		})
	}
}

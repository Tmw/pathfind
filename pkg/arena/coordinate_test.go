package arena

import (
	"reflect"
	"runtime"
	"testing"
)

func TestCoordinateDistanceTo(t *testing.T) {
	cases := []struct {
		c1, c2 Coordinate
		d      int
	}{
		{
			c1: NewCoordinate(1, 1),
			c2: NewCoordinate(2, 2),
			d:  2,
		},
		{
			c1: NewCoordinate(2, 2),
			c2: NewCoordinate(1, 1),
			d:  2,
		},
		{
			c1: NewCoordinate(5, 5),
			c2: NewCoordinate(-5, -5),
			d:  20,
		},
	}

	for _, c := range cases {
		actual := c.c1.DistanceTo(c.c2)
		if actual != c.d {
			t.Fatalf("expected distance between %+v and %+v to be %d but got %d", c.c1, c.c2, c.d, actual)
		}
	}
}

func TestCoordinateMove(t *testing.T) {
	cases := []struct {
		c  Coordinate
		fn func(Coordinate) Coordinate
		e  Coordinate
	}{
		{
			c:  NewCoordinate(1, 1),
			fn: Coordinate.North,
			e:  NewCoordinate(1, 0),
		},
		{
			c:  NewCoordinate(1, 1),
			fn: Coordinate.East,
			e:  NewCoordinate(2, 1),
		},
		{
			c:  NewCoordinate(1, 1),
			fn: Coordinate.West,
			e:  NewCoordinate(0, 1),
		},
		{
			c:  NewCoordinate(1, 1),
			fn: Coordinate.South,
			e:  NewCoordinate(1, 2),
		},
	}

	for _, c := range cases {
		actual := c.fn(c.c)
		if actual != c.e {
			fnName := runtime.FuncForPC(reflect.ValueOf(c.fn).Pointer()).Name()
			t.Fatalf("expected to get %+v when calling %+v on %+v, instead received %+v", c.e, fnName, c.c, actual)
		}
	}
}

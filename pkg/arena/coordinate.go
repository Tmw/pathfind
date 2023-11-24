package arena

import "math"

type Coordinate struct {
	x, y int
}

func NewCoordinate(x, y int) Coordinate {
	return Coordinate{x: x, y: y}
}

func (c Coordinate) DistanceTo(t Coordinate) int {
	n := (c.x - t.x) + (c.y - t.y)
	return int(math.Abs(float64(n)))
}

func (c Coordinate) North() Coordinate { return Coordinate{y: c.y - 1, x: c.x} }
func (c Coordinate) South() Coordinate { return Coordinate{y: c.y + 1, x: c.x} }
func (c Coordinate) West() Coordinate  { return Coordinate{y: c.y, x: c.x - 1} }
func (c Coordinate) East() Coordinate  { return Coordinate{y: c.y, x: c.x + 1} }

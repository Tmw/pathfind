package pathfind

// given a candidate that is marked as the finish,
// backtrace to the start and return the path from finish to start.
func backtrace[T comparable](c candidate[T]) []T {
	path := []T{}

	hop := c
	for hop.parent != nil {
		path = append(path, hop.coord)
		hop = *hop.parent
	}

	return path
}

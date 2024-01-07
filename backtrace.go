package pathfind

// given a candidate that is marked as the finish,
// backtrace to the start and return the path from finish to start.
func backtrace[T comparable](c candidate[T]) []T {
	path := []T{}

	hop := c
	for {
		path = append(path, hop.coord)
		if hop.parent == nil {
			break
		}
		hop = *hop.parent
	}

	return path
}

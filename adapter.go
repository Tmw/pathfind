package pathfind

type Adapter[T comparable] interface {
	// to return the next reachable steps from the given T
	Neighbours(T) []T

	// to return the estimated cost to reach the finish from the given T.
	// this function is currently only used when using the astar algorithm
	CostToFinish(T) int

	// indicating whether the given T is the finish or not
	IsFinish(T) bool
}

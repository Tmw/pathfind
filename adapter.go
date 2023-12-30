package pathfind

type Adapter[T comparable] interface {
	Neighbours(T) []T
	DistanceToFinish(T) int
	IsCellWalkable(T) bool
	IsCellFinish(T) bool
}

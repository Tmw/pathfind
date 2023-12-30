package pathfind

type FuncAdapter[T comparable] struct {
	NeighboursFn       func(T) []T
	DistanceToFinishFn func(T) int
	IsCellWalkableFn   func(T) bool
	IsCellFinishFn     func(T) bool
}

func (a *FuncAdapter[T]) Neighbours(c T) []T {
	return a.NeighboursFn(c)
}

func (a *FuncAdapter[T]) DistanceToFinish(c T) int {
	return a.DistanceToFinishFn(c)
}

func (a *FuncAdapter[T]) IsCellWalkable(c T) bool {
	return a.IsCellWalkableFn(c)
}

func (a *FuncAdapter[T]) IsCellFinish(c T) bool {
	return a.IsCellFinishFn(c)
}

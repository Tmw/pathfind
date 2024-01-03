package pathfind

type FuncAdapter[T comparable] struct {
	NeighboursFn   func(T) []T
	CostToFinishFn func(T) int
	IsFinishFn     func(T) bool
}

func (a *FuncAdapter[T]) Neighbours(c T) []T {
	return a.NeighboursFn(c)
}

func (a *FuncAdapter[T]) CostToFinish(c T) int {
	return a.CostToFinishFn(c)
}

func (a *FuncAdapter[T]) IsFinish(c T) bool {
	return a.IsFinishFn(c)
}

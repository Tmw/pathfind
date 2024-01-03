package pathfind

type SolveContext[T comparable] struct {
	MaxCost int

	Publish   func(Event)
	Adapter   func() Adapter[T]
	IsVisited func(T) bool
	Visit     func(T)
}

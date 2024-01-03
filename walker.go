package pathfind

type Walker[T comparable] interface {
	Walk(SolveContext[T]) []T
}

package pathfind

type candidate[T comparable] struct {
	coord  T
	parent *candidate[T]
	cost   int
}

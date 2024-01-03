package pathfind

type Algorithm int

const (
	AlgorithmBFS Algorithm = iota
	AlgorithmAStar
	AlgorithmDijkstra
)

type Solver[T comparable] struct {
	adapter  Adapter[T]
	eventlog []Event
	visited  map[T]struct{}

	// delegate algorithm
	walker Walker[T]

	// some runtime options
	MaxCost int
}

func (s *Solver[T]) isVisited(c T) bool {
	_, found := s.visited[c]
	return found
}

func (s *Solver[T]) visit(c T) {
	s.visited[c] = struct{}{}
}

func (s *Solver[T]) publish(e Event) {
	// Note: not threadsafe, yet?
	s.eventlog = append(s.eventlog, e)
}

func (s *Solver[T]) getAdapter() Adapter[T] {
	return s.adapter
}

func (s *Solver[T]) EventLog() []Event {
	return s.eventlog
}

func (s *Solver[T]) Walk() []T {
	return s.walker.Walk(SolveContext[T]{
		MaxCost:   s.MaxCost,
		Publish:   s.publish,
		Adapter:   s.getAdapter,
		IsVisited: s.isVisited,
		Visit:     s.visit,
	})
}

func makeWalker[T comparable](algorithm Algorithm, start T) Walker[T] {
	switch algorithm {
	case AlgorithmBFS:
		return newBFS[T](start)

	case AlgorithmAStar:
		return newAStar[T](start)

	default:
		return newAStar[T](start)
	}
}

func NewSolver[T comparable](algorithm Algorithm, start T, adapter Adapter[T]) Solver[T] {
	return Solver[T]{
		adapter:  adapter,
		walker:   makeWalker[T](algorithm, start),
		eventlog: []Event{},
		visited:  make(map[T]struct{}),
	}
}

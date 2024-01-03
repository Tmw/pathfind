package pathfind

type Event interface {
	event()
}

type EventCandidateAdded[T comparable] struct {
	CandidateID T
}

type EventCandidateVisited[T comparable] struct {
	CandidateID T
}

type EventFinishReached[T comparable] struct {
	Path []T
}

type EventUnsolvable struct{}
type EventMaxCostReached struct{}

func (e EventCandidateAdded[T]) event()   {}
func (e EventCandidateVisited[T]) event() {}
func (e EventFinishReached[T]) event()    {}
func (e EventUnsolvable) event()          {}
func (e EventMaxCostReached) event()      {}

package slice

// takes a slice and calls fn on each element of that slice,
// returns transformed slice.
func Map[C ~[]E, E comparable, O comparable](line C, fn func(E) O) []O {
	out := make([]O, len(line))
	for i := range line {
		out[i] = fn(line[i])
	}
	return out
}

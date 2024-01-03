package slice

// returns true if all elements of collection C satisfy the given predicate fn.
// Otherwise returns false.
func All[C comparable](c []C, fn func(C) bool) bool {
	for _, i := range c {
		if !fn(i) {
			return false
		}
	}

	return true
}

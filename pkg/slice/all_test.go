package slice

import (
	"testing"
)

func TestAll(t *testing.T) {
	t.Run("returns true if all conditions are met", func(t *testing.T) {
		evens := []int{2, 4, 6, 8, 10, 12}
		isEven := func(i int) bool {
			return i%2 == 0
		}

		if All(evens, isEven) != true {
			t.Error("expected true, received false")
		}
	})

	t.Run("returns false if not all conditions are met", func(t *testing.T) {
		evens := []int{2, 4, 6, 8, 10, 13}
		isEven := func(i int) bool {
			return i%2 == 0
		}

		if All(evens, isEven) != false {
			t.Error("expected false, received true")
		}
	})
}

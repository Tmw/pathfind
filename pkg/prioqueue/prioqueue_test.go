package prioqueue

import (
	"testing"
)

func assertEqual[T comparable](t *testing.T, a, b []T) {
	if len(a) != len(b) {
		t.Errorf("a and B lenghts differ: a = %d, b = %d", len(a), len(b))
	}

	for idx := range a {
		if a[idx] != b[idx] {
			t.Errorf("expected %v at index %d but got %v", a[idx], idx, b[idx])
		}
	}
}

func popAll[T comparable](q *Prioqueue[T]) []T {
	res := make([]T, q.Len())
	for i := 0; i < len(res); i++ {
		res[i] = q.Pop()
	}
	return res
}

func TestLen(t *testing.T) {
	q := New[string]()
	q.Push("first", 1)
	q.Push("second", 2)
	q.Push("third", 3)

	if q.Len() != 3 {
		t.Errorf("expected length of 3, got %d", q.Len())
	}
}

func TestPuhingAndPoppingInOrder(t *testing.T) {
	q := New[string]()

	items := []string{"first", "second", "third"}
	for idx, item := range items {
		q.Push(item, 10-idx)
	}

	out := popAll(q)
	assertEqual(t, out, []string{"third", "second", "first"})
}

func TestPeekItem(t *testing.T) {
	q := New[string]()
	q.Push("red", 10)
	q.Push("green", 10)
	q.Push("orange", 10)
	q.Push("pink", 10)

	got, want := q.PeekItem(2), "orange"

	if got != want {
		t.Errorf("expected %+v but got: %+v", want, got)
	}
}

func TestIndexfunc(t *testing.T) {
	q := New[string]()
	q.Push("red", 10)
	q.Push("green", 10)
	q.Push("orange", 10)
	q.Push("pink", 10)

	t.Run("when found", func(t *testing.T) {
		predicate := func(item string) bool {
			return item == "orange"
		}

		got, want := q.IndexFunc(predicate), 2
		if got != want {
			t.Errorf("expected %d but got: %d", want, got)
		}
	})

	t.Run("when not found", func(t *testing.T) {
		predicate := func(item string) bool {
			return item == "yellow"
		}

		got, want := q.IndexFunc(predicate), -1
		if got != want {
			t.Errorf("expected %d but got: %d", want, got)
		}
	})
}

func TestPriorityOfItem(t *testing.T) {
	q := New[string]()
	q.Push("red", 10)
	q.Push("green", 20)
	q.Push("orange", 30)
	q.Push("pink", 40)

	if q.PriorityOfItem(2) != 30 {
		t.Errorf("expected priority of item with index 2 to be 30, got: %d", q.PriorityOfItem(2))
	}
}

func TestUpdateAtIndex(t *testing.T) {
	q := New[string]()
	q.Push("red", 10)
	q.Push("green", 20)
	q.Push("orange", 30)
	q.Push("pink", 40)

	q.UpdateAtIndex(1, "blue", 50)
	actual, expected := popAll(q), []string{"red", "orange", "pink", "blue"}
	assertEqual(t, expected, actual)
}

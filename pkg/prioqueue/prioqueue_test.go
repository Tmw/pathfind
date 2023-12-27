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

	out := []string{}
	for i := 0; i < len(items); i++ {
		out = append(out, q.Pop())
	}

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

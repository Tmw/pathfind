package queue

import (
	"reflect"
	"testing"
)

func TestLen(t *testing.T) {
	q := New[int](1, 2, 3)

	if q.Len() != 3 {
		t.Errorf("expected length of 3, got %d", q.Len())
	}
}

func TestEmpty(t *testing.T) {
	q := New[int](1, 2, 3)

	if q.Empty() != false {
		t.Error("expected queue to be not empty")
	}

	q.Pop()
	q.Pop()
	q.Pop()

	if q.Empty() != true {
		t.Error("expected queue to be empty")
	}
}

func TestPushAndPop(t *testing.T) {
	q := New[int](1, 2, 3)
	q.Push(4, 5, 6)

	expected := []int{1, 2, 3, 4, 5, 6}
	actual := []int{}

	for !q.Empty() {
		actual = append(actual, q.Pop())
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%+v does not equal %+v", actual, expected)
	}
}

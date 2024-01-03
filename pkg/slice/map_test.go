package slice

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	expected := []int{2, 4, 6, 8, 10}

	double := func(i int) int {
		return i * 2
	}

	actual := Map(nums, double)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %+v to equal %+v", actual, expected)
	}
}

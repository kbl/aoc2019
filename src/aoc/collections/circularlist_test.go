package collections

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	cl := NewCircularList()
	if cl.size != 0 {
		t.Errorf("NewCircularList().Size() = %d, want 0", cl.Size())
	}

	cl.Add(1)
	cl.Add(2)

	if cl.size != 2 {
		t.Errorf("cl.Add().Size() = %d, want 0", cl.Size())
	}
}

func TestToSlice(t *testing.T) {
	cl := NewCircularList()
	cl.Add(1)
	cl.Add(2)
	cl.Add(3)
	cl.Add(4)

	got := cl.ToSlice(Forward)
	expected := []int{4, 1, 2, 3}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("cl.ToSlice(Forward) = %v, want %v", got, expected)
	}

	got = cl.ToSlice(Backward)
	expected = []int{4, 3, 2, 1}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("cl.ToSlice(Backward) = %v, want %v", got, expected)
	}
}

func TestShift(t *testing.T) {
	cl := NewCircularList()
	cl.Add(1)
	cl.Add(2)
	cl.Add(3)
	cl.Add(4)

	cl.LShift()
	cl.LShift()

	got := cl.ToSlice(Forward)
	expected := []int{2, 3, 4, 1}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("cl.ToSlice(Forward) = %v, want %v", got, expected)
	}

	got = cl.ToSlice(Backward)
	expected = []int{2, 1, 4, 3}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("cl.ToSlice(Backward) = %v, want %v", got, expected)
	}
}

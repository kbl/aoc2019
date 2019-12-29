package collections

import (
	"reflect"
	"testing"
)

func TestAppend(t *testing.T) {
	d := NewDeque()
	if d.size != 0 {
		t.Errorf("NewDeque().Size() = %d, want 0", d.Size())
	}

	d.Append(1)
	d.Append(2)

	if d.size != 2 {
		t.Errorf("d.Append().Size() = %d, want 0", d.Size())
	}
}

func TestToSlice(t *testing.T) {
	d := NewDeque()
	d.Append(1)
	d.Append(2)
	d.Append(3)
	d.Append(4)

	got := d.ToSlice(Forward)
	expected := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("d.ToSlice(Forward) = %v, want %v", got, expected)
	}

	got = d.ToSlice(Backward)
	expected = []int{4, 3, 2, 1}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("d.ToSlice(Backward) = %v, want %v", got, expected)
	}
}

func TestRotate(t *testing.T) {
	d := NewDeque()
	d.Append(1)
	d.Append(2)
	d.Append(3)
	d.Append(4)

	d.Rotate(2)

	got := d.ToSlice(Forward)
	expected := []int{3, 4, 1, 2}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("d.ToSlice(Forward) = %v, want %v", got, expected)
	}

	got = d.ToSlice(Backward)
	expected = []int{2, 1, 4, 3}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("d.ToSlice(Backward) = %v, want %v", got, expected)
	}

	d.Rotate(-2)

	got = d.ToSlice(Forward)
	expected = []int{1, 2, 3, 4}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("d.ToSlice(Forward) = %v, want %v", got, expected)
	}

	got = d.ToSlice(Backward)
	expected = []int{4, 3, 2, 1}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("d.ToSlice(Backward) = %v, want %v", got, expected)
	}
}

func TestPop(t *testing.T) {
	d := NewDeque()
	d.Append(1)
	d.Append(2)
	d.Append(3)

	got, _ := d.Pop()
	expected := 3

	if got != expected {
		t.Errorf("d.Pop() = %v, want %v", got, expected)
	}

	got, _ = d.Pop()
	expected = 2

	if got != expected {
		t.Errorf("d.Pop() = %v, want %v", got, expected)
	}

	got, _ = d.Pop()
	expected = 1

	if got != expected {
		t.Errorf("d.Pop() = %v, want %v", got, expected)
	}
}

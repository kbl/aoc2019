package main

import (
	"reflect"
	"strings"
	"testing"
)

type testCase struct {
	input    int
	expected int
}

var testCases = []testCase{}

func TestNewDeck(t *testing.T) {
	d := NewDeck(10)
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	got := d.Content()

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("NewDeck(10) = %d, want %d", got, expected)
	}
}

func TestDealInto(t *testing.T) {
	d := NewDeck(10)
	expected := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	d.Deal()
	got := d.Content()

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("NewDeck(10).Deal() = %d, want %d", got, expected)
	}
}

func TestCut(t *testing.T) {
	d := NewDeck(10)
	expected := []int{3, 4, 5, 6, 7, 8, 9, 0, 1, 2}
	d.Cut(3)
	got := d.Content()

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("NewDeck(10).Cut(3) = %d, want %d", got, expected)
	}

	d = NewDeck(10)
	expected = []int{6, 7, 8, 9, 0, 1, 2, 3, 4, 5}
	d.Cut(-4)
	got = d.Content()

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("NewDeck(10).Cut(-4) = %d, want %d", got, expected)
	}
}

func TestIncrement(t *testing.T) {
	d := NewDeck(10)
	expected := []int{0, 7, 4, 1, 8, 5, 2, 9, 6, 3}
	d.Increment(3)
	got := d.Content()

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("NewDeck(10).Increment(3) = %d, want %d", got, expected)
	}
}

func TestPosition(t *testing.T) {
	d := NewDeck(10)
	expected := 3
	got := d.Position(3)
	if got != expected {
		t.Errorf("NewDeck(10).Position(3) = %d, want %d", got, expected)
	}
}

func TestShuffle(t *testing.T) {
	instructions := strings.Split(`deal with increment 7
deal into new stack
deal into new stack`, "\n")
	d := NewDeck(10)
	d.Shuffle(instructions)
	expected := []int{0, 3, 6, 9, 2, 5, 8, 1, 4, 7}
	got := d.Content()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("NewDeck(10).Shuffle() = %d, want %d", got, expected)
	}

	instructions = strings.Split(`cut 6
deal with increment 7
deal into new stack`, "\n")
	d = NewDeck(10)
	d.Shuffle(instructions)
	expected = []int{3, 0, 7, 4, 1, 8, 5, 2, 9, 6}
	got = d.Content()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("NewDeck(10).Shuffle() = %d, want %d", got, expected)
	}

	instructions = strings.Split(`deal with increment 7
deal with increment 9
cut -2`, "\n")
	d = NewDeck(10)
	d.Shuffle(instructions)
	expected = []int{6, 3, 0, 7, 4, 1, 8, 5, 2, 9}
	got = d.Content()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("NewDeck(10).Shuffle() = %d, want %d", got, expected)
	}

	instructions = strings.Split(`deal into new stack
cut -2
deal with increment 7
cut 8
cut -4
deal with increment 7
cut 3
deal with increment 9
deal with increment 3
cut -1`, "\n")
	d = NewDeck(10)
	d.Shuffle(instructions)
	expected = []int{9, 2, 5, 8, 1, 4, 7, 0, 3, 6}
	got = d.Content()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("NewDeck(10).Shuffle() = %d, want %d", got, expected)
	}
}

func TestFunctions(t *testing.T) {
	tcs := [][]string{
		{"deal with increment 7"},
		{"deal into new stack"},
		{
			"deal into new stack",
			"deal into new stack",
		},
		{
			"deal into new stack",
			"cut -2",
			"cut 8",
			"deal into new stack",
			"cut -4",
			"cut 3",
			"cut -1",
		},
	}
	size := 10007

	for _, instructions := range tcs {
		d := NewDeck(size)
		d.Shuffle(instructions)
		expected := d.Content()

		for index, finalValue := range expected {
			got := trackFunctions(instructions, index, size)
			if got != finalValue {
				t.Errorf("trackFunctions(%v, %d, %d) = %d, want %d", instructions, index, size, got, finalValue)
			}
		}
	}
}

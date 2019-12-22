package main

import (
	"reflect"
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

	d = NewDeck(10)
	expected = []int{6, 7, 8, 9, 0, 1, 2, 3, 4, 5}
	d.Cut(-4)
	got = d.Content()

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("NewDeck(10).Cut(-4) = %d, want %d", got, expected)
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

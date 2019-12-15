package main

import (
	"reflect"
	"testing"
)

/*findling path from {0 0} to {0 1}
{{1 0} [4]}
{{0 1} [1]}
move: ^
 moving: ^ from {0 0}
 found: # (0) at {0 0} -> {0 1}
?????
??#??
? *#?
??#??
?????
{0 0}

findling path from {0 0} to {-1 1}*/

func TestSomething(t *testing.T) {
	d := &Droid{
		grid: Grid{
			{0, 0}:  moved,
			{0, 1}:  wall,
			{1, 0}:  wall,
			{0, -1}: wall,
			{-1, 0}: moved,
		},
	}

	got := d.moveSequence(Cord{-1, 1})
	expected := []move{3, 1}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("d.moveSequence() = %v, want %v", got, expected)
	}
}

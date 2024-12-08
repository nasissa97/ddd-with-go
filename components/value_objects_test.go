package components_test

import (
	"testing"

	"github.com/nasissa97/ddd-with-go/components"
)

func Test_Point(t *testing.T) {
	a := components.NewPoint(1, 1)
	b := components.NewPoint(1, 1)
	c := components.NewPoint(3, 3)
	if a != b {
		t.Fatalf("a and b were not equal")
	}

	if a == c {
		t.Fatalf("a and c were equal")
	}
}

func Test_Moves(t *testing.T) {
	start := components.NewPoint(1, 1)
	end := components.NewPoint(2, 3)
	start = components.Move(start, components.DirectionNorth)
	start = components.Move(start, components.DirectionNorth)
	start = components.Move(start, components.DirectionEast)

	if start != end {
		t.Fatalf("start should have moved to end")
	}
}

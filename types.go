package main

import (
	"fmt"
)

//    /        /        /
// vertex--edge--vertex
//   |      |       |   /
//  edge---side---edge
//   |      |       |   /
// vertex--edge--vertex
//
//

type box interface {
	Move(direction _Direction)
	Get() string
}

// Side ...
type Side struct {
	colors Color
}

// Move description
func (s *Side) Move(
	direction _Direction,
) {
	fmt.Printf("\nmove to %v\n", direction)

	for ii := 0; ii < dimension; ii++ {
		if direction = direction & _Direction(dimensionMask); direction == 0 {
			fmt.Printf("not direction for move %v\n", direction)
			return
		}
	}

	if direction&s.colors.orientation != 0 ||
		direction.Reverse()&s.colors.orientation != 0 {
		fmt.Printf("side box cann't move to this posistion curent %v move to %v\n", s.colors, direction)
		// panic(fmt.Errorf("side box cann't move to this posistion curent %v move to %v", s.colors, direction))
	}

	fmt.Printf("OLD color: %v\n", s.colors)

	s.colors.orientation = direction
	fmt.Printf("NEW color: %v \n", s.colors)
}

// Get description
func (s *Side) Get() string {
	return s.colors.String()
}

// Edge ...
type Edge struct {
}

// Move description
func (e *Edge) Move(
	direction _Direction,
) {
	fmt.Printf("move to %08b\n", direction)
}

// Get description
func (e *Edge) Get() string {
	return ""
}

// Vertex ...
type Vertex struct {
}

// Move description
func (v *Vertex) Move(
	direction int,
	increase bool,
) {
	fmt.Printf("move to %08b\n", direction)
}

// Get description
func (v *Vertex) Get() string {
	return "⸢⸣⸤⸥⸠⸡"
}

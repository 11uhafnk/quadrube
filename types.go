package main

import (
	"fmt"
	"strings"
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
	Move(plane _Plane)
	Get() string
}

// Side ...
// count Color must be 1
type Side struct {
	colors Color
}

// Move for side box need only first direction
func (s *Side) Move(
	plane _Plane,
) {
	fmt.Printf("\nmove to %v\n", plane)
	direction := plane[0]
	if !direction.Check() {
		fmt.Printf("wrong direction for move %v\n", direction)
		return
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
// count Color must be equal the dimension-1
type Edge struct {
	colors []Color
}

// Move description
func (e *Edge) Move(
	plane _Plane,
) {
	var status string

	fmt.Print("move to ", plane, "\t\t")
	if !plane.Check() {
		fmt.Printf("wrong rotate plane %v\n", plane)
		return
	}

	if len(e.colors) != dimension-1 {
		panic(fmt.Errorf("edge haven't %d sides", dimension-1))
	}

	fmt.Print("OLD:", e.Get(), "\t\t")

	index0, index1, canMove := checkMove(plane, e.colors)

	if !canMove || (index0 != -1 && index1 == -1) {
		status = fmt.Sprintf("edge box cann't move to this posistion curent %v move to %v\t", e.Get(), plane)
	} else {
		status = "OK"
	}

	if index0 != -1 && index1 != -1 {
		e.colors[index0].orientation, e.colors[index1].orientation = e.colors[index1].orientation, plane[0]
	} else if index1 != -1 { // outer
		e.colors[index1].orientation = plane[0]
	} else {
		// nothing
		// need change sides only at the rotating plane

	}

	fmt.Println("NEW:", e.Get(), "\t\t", status)

}

// Get description
func (e *Edge) Get() string {
	builder := strings.Builder{}
	builder.WriteString("[ ")
	for ii := range e.colors {
		if ii != 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(e.colors[ii].String())
	}
	builder.WriteString("]")

	return builder.String()
}

// Vertex ...
// count Color must be equal the dimension-1
type Vertex struct {
	colors []Color
}

// Move description
func (v *Vertex) Move(
	plane _Plane,
) {
	var status string

	fmt.Print("move to ", plane, "\t\t")
	if !plane.Check() {
		fmt.Printf("wrong rotate plane %v\n", plane)
		return
	}

	if len(v.colors) != dimension {
		panic(fmt.Errorf("edge haven't %v sides", dimension))
	}

	fmt.Print("OLD:", v.Get(), "\t\t")

	index0, index1, canMove := checkMove(plane, v.colors)

	if !canMove || (index0 != -1 && index1 == -1) {
		status = fmt.Sprintf("edge box cann't move to this posistion curent %v move to %v\t", v.Get(), plane)
	} else {
		status = "OK"
	}

	if index0 != -1 && index1 != -1 {
		v.colors[index0].orientation, v.colors[index1].orientation = v.colors[index1].orientation, plane[0]
	} else if index1 != -1 { // outer
		v.colors[index1].orientation = plane[0]
	} else {
		// nothing
		// need change sides only at the rotating plane

	}

	fmt.Println("NEW:", v.Get(), "\t\t", status)

}

// Get description
func (v *Vertex) Get() string {
	builder := strings.Builder{}
	builder.WriteString("[ ")
	for ii := range v.colors {
		if ii != 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(v.colors[ii].String())
	}
	builder.WriteString("]")

	return builder.String()
}

func checkMove(
	plane _Plane,
	colors []Color,
) (
	index0 int,
	index1 int,
	canMove bool,
) {
	index0, index1 = -1, -1
	notP0 := plane[0].Reverse()
	notP1 := plane[1].Reverse()

	for ii := range colors {

		//  /   /   /
		// *---*---*
		// |   |   | /        ^         ^ -->
		// *---*---X      -->/    or   /       isn't correct
		// |   |   | /
		// *---*---*
		//
		// если блок показывает цыет в какую-то сторону, то это крайнее положение
		// нельзя двигать бокс дальше в этом направлении
		if colors[ii].orientation == plane[0] ||
			colors[ii].orientation == plane[1] {
			return index0, index1, false
		}

		//  /   /   /   inner                /   /   /  outer          Y
		// *---*---*                        *---X---*                  ^   ^ Z
		// |   |   | /        ^             |   |   | /       ^        |  /
		// X---*---*      -->/      else    *---*---*     -->/         | /
		// |   |   | /                      |   |   | /                |/
		// *---*---*                        *---X---*                  *------> X
		//
		// X+ Z+                               X+ Z+
		// !X , !Z -> !Z , X
		//  ^    ^ inner1
		//  inner0   because move order
		if colors[ii].orientation == notP0 {
			index0 = ii
		}
		if colors[ii].orientation == notP1 {
			index1 = ii
		}
	}

	return index0, index1, true
}

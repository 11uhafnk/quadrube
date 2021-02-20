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
	Move(plane _Plane) error
	Get() string
}

// Side ...
// count Color must be 1
type Side struct {
	colors Color
}

// Move for side box need only first direction
func (box *Side) Move(
	plane _Plane,
) (
	err error,
) {
	fmt.Printf("move to %v\t\t", plane)
	if err = plane.Check(); err != nil {
		return fmt.Errorf("%w: plane %v", err, plane)
	}

	index0, index1, err := checkMove(plane, []Color{box.colors})
	if err != nil {
		err = fmt.Errorf("%v %w: curent %v move to %v", "side", err, box.colors, plane)
	} else if index0 != -1 {
		err = fmt.Errorf("%v %w: curent %v move to %v", "side", ErrMove, box.colors, plane)
	}

	fmt.Printf("OLD color: %v\t\t", box.Get())

	if index1 != -1 {
		box.colors.orientation = plane[0]

	}

	fmt.Printf("NEW color: %v \t\t%v\n", box.Get(), err)
	return err
}

// Get description
func (box *Side) Get() string {
	return box.colors.String()
}

// Edge ...
// count Color must be equal the dimension-1
type Edge struct {
	colors []Color
}

// Move description
func (box *Edge) Move(
	plane _Plane,
) (
	err error,
) {

	fmt.Print("move to ", plane, "\t\t")
	if err = plane.Check(); err != nil {
		return fmt.Errorf("%w: plane %v", err, plane)
	}

	if len(box.colors) != dimension-1 {
		panic(fmt.Errorf("edge haven't %d sides", dimension-1))
	}

	fmt.Print("OLD:", box.Get(), "\t\t")

	index0, index1, err := checkMove(plane, box.colors)
	if err != nil {
		err = fmt.Errorf("%v %w: curent %v move to %v", "edge", err, box.colors, plane)
	} else if index0 != -1 && index1 == -1 {
		err = fmt.Errorf("%v %w: curent %v move to %v", "edge", ErrMove, box.colors, plane)
	}
	err = fmt.Errorf("%d, %d, %w", index0, index1, err)

	if index0 != -1 && index1 != -1 {
		box.colors[index0].orientation, box.colors[index1].orientation = box.colors[index1].orientation, plane[0]
	} else if index1 != -1 { // outer
		box.colors[index1].orientation = plane[0]
	}

	fmt.Println("NEW:", box.Get(), "\t", index0, index1, "\t", err)
	return err

}

// Get description
func (box *Edge) Get() string {
	builder := strings.Builder{}
	builder.WriteString("[ ")
	for ii := range box.colors {
		if ii != 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(box.colors[ii].String())
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
func (box *Vertex) Move(
	plane _Plane,
) (
	err error,
) {

	fmt.Print("move to ", plane, "\t\t")
	if err = plane.Check(); err != nil {
		err := fmt.Errorf("%w: plane %v", err, plane)
		fmt.Println(err)
		return err
	}

	if len(box.colors) != dimension {
		panic(fmt.Errorf("edge haven't %v sides", dimension))
	}

	fmt.Print("OLD:", box.Get(), "\t\t")

	index0, index1, err := checkMove(plane, box.colors)
	if err != nil {
		err = fmt.Errorf("%v %w: curent %v move to %v", "vertex", err, box.colors, plane)
	} else if index0 != -1 && index1 == -1 {
		err = fmt.Errorf("%v %w: curent %v move to %v", "vertex", ErrMove, box.colors, plane)
	}

	if index0 != -1 && index1 != -1 {
		box.colors[index0].orientation, box.colors[index1].orientation = box.colors[index1].orientation, plane[0]
	} else if index1 != -1 { // outer
		box.colors[index1].orientation = plane[0]
	}

	fmt.Println("NEW:", box.Get(), "\t", index0, index1, "\t", err)
	return err
}

// Get description
func (box *Vertex) Get() string {
	builder := strings.Builder{}
	builder.WriteString("[ ")
	for ii := range box.colors {
		if ii != 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(box.colors[ii].String())
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
	err error,
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
			return index0, index1, ErrMove
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

	return index0, index1, nil
}

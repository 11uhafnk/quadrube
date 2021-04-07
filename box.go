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

// Box ...
type Box interface {
	Move(plane _Plane) error
	Get() string
}

// Empty no touchSides
type Empty struct{}

// Move ...
func (box *Empty) Move(
	plane _Plane,
) (
	err error,
) {
	return nil
}

func (box *Empty) Get() string {
	mask := fmt.Sprintf("%%%ds", _Dimension*4+2)
	return fmt.Sprintf(mask, " ")
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
	defer debugPrintln("")
	debugPrintf("move to %v\t\t", plane)
	if err = plane.Check(); err != nil {
		return fmt.Errorf("%w: plane %v", err, plane)
	}

	index0, index1, err := checkMove(plane, []Color{box.colors})
	if err != nil {
		err = fmt.Errorf("%v %w: curent %v move to %v", "side", err, box.colors, plane)
	} else if index0 != -1 {
		err = fmt.Errorf("%v %w: curent %v move to %v", "side", ErrMove, box.colors, plane)
	}

	debugPrintf("OLD color: %v\t\t", box.Get())

	if index1 != -1 {
		box.colors.orientation = plane[0]

	}

	debugPrintf("NEW color: %v \t\t%v", box.Get(), err)
	return err
}

// Get description
func (box *Side) Get() string {

	mask := fmt.Sprintf("%%%ds%%s%%%ds", (_Dimension-1)*2+2, (_Dimension-1)*2+2)
	return fmt.Sprintf(mask, " ", box.colors.String(), " ")
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
	defer debugPrintln("")
	debugPrint("move to ", plane, "\t\t")
	if err = plane.Check(); err != nil {
		return fmt.Errorf("%w: plane %v", err, plane)
	}

	if len(box.colors) != _Dimension-1 {
		panic(fmt.Errorf("edge haven't %d sides", _Dimension-1))
	}

	debugPrint("OLD:", box.Get(), "\t\t")

	index0, index1, err := checkMove(plane, box.colors)
	if err != nil {
		err = fmt.Errorf("%v %w: curent %v move to %v", "edge", err, box.colors, plane)
		return
	} else if index0 != -1 && index1 == -1 {
		err = fmt.Errorf("%v %w: curent %v move to %v", "edge", ErrMove, box.colors, plane)
		return
	}

	if index0 != -1 && index1 != -1 {
		box.colors[index0].orientation, box.colors[index1].orientation = box.colors[index1].orientation, plane[0]
	} else if index1 != -1 { // outer
		box.colors[index1].orientation = plane[0]
	}

	debugPrint("NEW:", box.Get(), "\t", index0, index1, "\t", err)
	return err

}

// Get description
func (box *Edge) Get() string {
	builder := strings.Builder{}
	for ii := 0; ii < _Dimension-len(box.colors); ii++ {
		builder.WriteString("  ")
	}
	builder.WriteString("[ ")
	for ii := range box.colors {
		if ii != 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(box.colors[ii].String())
	}
	builder.WriteString("] ")

	for ii := 0; ii < _Dimension-len(box.colors); ii++ {
		builder.WriteString("  ")
	}

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
	defer debugPrintln("")
	debugPrint("move to ", plane, "\t\t")
	if err = plane.Check(); err != nil {
		err := fmt.Errorf("%w: plane %v", err, plane)
		debugPrint(err)
		return err
	}

	if len(box.colors) != _Dimension {
		panic(fmt.Errorf("edge haven't %v sides", _Dimension))
	}

	debugPrint("OLD:", box.Get(), "\t\t")

	index0, index1, err := checkMove(plane, box.colors)
	if err != nil {
		err = fmt.Errorf("%v %w: curent %v move to %v", "vertex", err, box.colors, plane)
		return
	}

	if index0 != -1 && index1 != -1 {
		box.colors[index0].orientation, box.colors[index1].orientation = box.colors[index1].orientation, plane[0]
	}

	debugPrint("NEW:", box.Get(), "\t", index0, index1, "\t", err)
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
	builder.WriteString("] ")

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

	debugPrintln(index0, index1, plane, colors)
	return index0, index1, nil
}

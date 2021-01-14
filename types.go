package main

import (
	"fmt"
	"strconv"
)

// 3    1    2
//  V   Y   Z
//   ^  ^  ^
//    \ | /
//     \|/
//      ------> X - 0

var (
	dimension = 3

	size = 2
)

type color int

func (c color) String() string {
	letter := strconv.Itoa(int(c))
	return "\033[0;3" + letter + "m" + letter + "\033[0;39m"
}

const (
	positionX = 1 << iota
	positionNX
	positionY
	positionNY
	positionZ
	positionNZ
	positionV
	positionNV
)

const (
	maskPositionX = 3 << (2 * iota)
	maskPositionY
	maskPositionZ
	maskPositionV
)

// Colors
const (
	ColorBlack color = iota
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
	colorEnd
	colorBase
)

// Color ...
type Color struct {
	color       color
	orientation int
}

type box interface {
	// GetCountColors() int
	Move(direction int, increase bool)
	Get() string
}

// Side ...
type Side struct {
	colors Color
}

// Move description
func (s *Side) Move(
	direction int,
	increase bool,
) {
	inc := "increase"
	if !increase {
		inc = "decrease"
	}
	fmt.Printf("move to %d to %s\n", direction, inc)
	fmt.Printf("OLD color: %v\n", s.colors)

	fmt.Printf("NEW color: %v \n", s.colors)
}

// Get description
func (s *Side) Get() string {
	return fmt.Sprintf("color: %v ", s.colors)
}

// Edge ...
type Edge struct {
}

// Move description
func (e *Edge) Move(
	direction int,
	increase bool,
) {
	inc := "increase"
	if !increase {
		inc = "decrease"
	}
	fmt.Printf("move to %d to %s", direction, inc)
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
	inc := "increase"
	if !increase {
		inc = "decrease"
	}
	fmt.Printf("move to %d to %s", direction, inc)
}

// Get description
func (v *Vertex) Get() string {
	return "⸢⸣⸤⸥⸠⸡"
}

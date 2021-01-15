package main

import (
	"strconv"
)

type color int

func (c color) String() string {
	letter := strconv.Itoa(int(c))
	return "\033[0;3" + letter + "m" + letter + "\033[0;39m"
}

func (c color) Brush() string {
	letter := strconv.Itoa(int(c))
	return "\033[0;3" + letter + "m"
}
func (c color) CLear() string {
	return "\033[0;39m"
}

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
	orientation _Direction
}

func (c Color) String() string {
	return c.color.Brush() + c.orientation.String() + c.color.CLear()
}

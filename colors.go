package main

import (
	"strconv"
)

type _Color int

func (c _Color) String() string {
	letter := strconv.Itoa(int(c))
	return "\033[0;3" + letter + "m" + letter + "\033[0;39m"
}

func (c _Color) Brush() string {
	letter := strconv.Itoa(int(c))
	return "\033[0;3" + letter + "m"
}
func (c _Color) CLear() string {
	return "\033[0;39m"
}

// Colors
const (
	ColorBlack _Color = iota
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
	color       _Color
	orientation _Direction
}

func (c Color) String() string {
	return c.color.Brush() + c.orientation.String() + c.color.CLear()
}

package main

import (
	"strconv"
)

// _Color для консольного вывода цветов
type _Color int

// String вернуть цвет в цвете
func (c _Color) String() string {
	letter := strconv.Itoa(int(c))
	return "\033[0;3" + letter + "m" + letter + "\033[0;39m"
}

// Brush включить у консоли текущий цвет
func (c _Color) Brush() string {
	letter := strconv.Itoa(int(c))
	return "\033[0;3" + letter + "m"
}

// CLear вернуть сброс к стандартному цвету консоли
func (c _Color) CLear() string {
	return "\033[0;39m"
}

// Colors текущие доступные цвета
const (
	ColorBlack   _Color = iota // 0
	ColorRed                   // 1
	ColorGreen                 // 2
	ColorYellow                // 3
	ColorBlue                  // 4
	ColorMagenta               // 5
	ColorCyan                  // 6
	ColorWhite                 // 7
	colorEnd                   //
	colorBase                  // 9
)

// Color цветовая поверхность на блоке
// имеет цвет и ориентацию в пространстве но только в одной координате
type Color struct {
	color       _Color
	orientation _Direction
}

// String возвращает направление куда повернута поверхность
// в цвете этой поверхности
func (c Color) String() string {
	return c.color.Brush() + c.orientation.String() + c.color.CLear()
}

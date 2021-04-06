package main

import (
	"fmt"
	"math"
)

// Cube ...
type Cube struct {
	// it as [][][][]box
	//        ^ deep as dimention
	arr []Box
}

func InitCube(
	dimension int,
	size int,
) (
	cube *Cube,
) {
	arrLen := int(math.Pow(float64(size), float64(dimension)))
	_Size = size
	_Dimension = dimension
	for ii := 0; ii < _Dimension; ii++ {
		_DimensionMask |= int(directions[ii])
	}

	initColors := make(map[_Direction]_Color, _Dimension*2)
	for ii := 0; ii < _Dimension; ii++ {
		initColors[directionX<<(ii*2)] = _Color(ii * 2)
		initColors[directionNX<<(ii*2)] = _Color(ii*2 + 1)
	}

	c := Cube{
		arr: make([]Box, 0, arrLen),
	}

	for ii := 0; ii < arrLen; ii++ {
		touchSides, dirs := getTouchSides(parsePosition(ii))
		switch touchSides {
		case 1:
			c.arr[ii] = &Side{colors: Color{color: initColors[dirs], orientation: dirs}}
		case _Dimension - 1:
			colors := make([]Color, 0, _Dimension-1)
			for _, dir := range dirs.Split() {
				colors = append(colors, Color{
					color:       initColors[dir],
					orientation: dir,
				})
			}
			c.arr[ii] = &Edge{
				colors: colors,
			}
		case _Dimension:
			colors := make([]Color, 0, _Dimension-1)
			for _, dir := range dirs.Split() {
				colors = append(colors, Color{
					color:       initColors[dir],
					orientation: dir,
				})
			}
			c.arr[ii] = &Vertex{
				colors: colors,
			}
		}
	}

	return &c
}

// Rotate description
func (c *Cube) Move(
	position []int,
	direction _Plane,
) {
	index0, index1 := findIndexes(direction)

	var coords []int
	copy(coords, position)

	for ii := 0; ii < _Size; ii++ {
		for jj := 0; jj < _Size; jj++ {
			coords[index0] = ii
			coords[index1] = jj
			fmt.Printf("%v %d\t", coords, getPosition(coords...))
		}
		fmt.Println("")
	}
}

func findIndexes(
	p _Plane,
) (
	index0 int,
	index1 int,
) {
	for ii := 0; ii < _Dimension; ii++ {
		if (maskDirectionX<<(ii*2))&p[0] != 0 {
			index0 = ii
		}
		if (maskDirectionX<<(ii*2))&p[1] != 0 {
			index1 = ii
		}
	}
	return
}

// getPosition( X Y Z V ...)
func getPosition(
	coords ...int,
) (
	coord int,
) {
	if len(coords) != _Dimension {
		return -1
	}

	for ii := 0; ii < _Dimension; ii++ {
		if coords[ii] >= _Dimension {
			return -2
		}
		coord += int(math.Pow(float64(_Size), float64(ii))) * coords[ii]
	}

	return coord
}

func parsePosition(
	coord int,
) (
	coords []int,
) {
	if coord < 0 || coord >= int(math.Pow(float64(_Size), float64(_Dimension))) {
		return []int{}
	}

	coords = make([]int, _Dimension)

	for ii := 0; ii < _Dimension; ii++ {
		coords[ii] = coord % _Size
		coord /= _Size
	}
	return
}

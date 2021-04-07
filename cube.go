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
		color := _Color(ii * 2)
		if color == ColorBlack {
			color = colorBase
		} else if color == colorEnd {
			color = ColorBlack
		}
		initColors[directionX<<(ii*2)] = color
		initColors[directionNX<<(ii*2)] = _Color(ii*2 + 1)
	}

	c := Cube{
		arr: make([]Box, arrLen),
	}

	for ii := 0; ii < arrLen; ii++ {
		touchSides, dirs := getTouchSides(parsePosition(ii))
		switch touchSides {
		case 0:
			c.arr[ii] = &Empty{}
		case 1:
			c.arr[ii] = &Side{colors: Color{color: initColors[dirs], orientation: dirs}}
		case _Dimension:
			colors := make([]Color, 0, _Dimension)
			for _, dir := range dirs.Split() {
				colors = append(colors, Color{
					color:       initColors[dir],
					orientation: dir,
				})
			}
			c.arr[ii] = &Vertex{
				colors: colors,
			}
		default:
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
		}
	}

	return &c
}

func (c *Cube) Print() {
	for ii := 0; ii < len(c.arr); ii++ {
		for jj := 1; jj < _Dimension; jj++ {
			qwe := int(math.Pow(float64(_Size), float64(jj)))
			if ii%qwe == 0 {
				fmt.Println("")
			}
		}
		fmt.Print(c.arr[ii].Get())
	}
	fmt.Println("")
}

// Rotate description
func (c *Cube) Move(
	position []int,
	direction _Plane,
) error {
	if len(position) != _Dimension {
		return ErrWrongDimention
	}

	index0, index1 := findIndexes(direction)

	coords := make([]int, len(position))
	copy(coords, position)

	moved := make(map[int]Box)

	for ii := 0; ii < _Size; ii++ {
		for jj := 0; jj < _Size; jj++ {
			coords[index0] = ii
			coords[index1] = jj
			oldPos := getPosition(coords...)
			dir := direction.ReDirection(coords)
			fmt.Printf("%v %d %v\t", coords, oldPos, dir)

			err := c.arr[oldPos].Move(direction.ReDirection(coords))
			if err != nil {
				return err
			}

			// moved[oldPos] = c.arr[oldPos]
		}
		fmt.Println("")
	}

	for pos := range moved {
		fmt.Printf("%+v\n", moved[pos].Get())
	}

	return nil
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
		if coords[ii] >= _Size {
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

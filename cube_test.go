package main

import (
	"math"
	"reflect"
	"testing"
)

func TestCubeGetPosition(t *testing.T) {

	check := func(coords []int, want int) {
		have := getPosition(coords...)
		if have != want {
			t.Errorf("WANT %v HAVE %v", want, have)
		}
	}

	_Dimension = 2
	check([]int{0}, -1)
	check([]int{0, 0}, 0)
	check([]int{0, 0, 0}, -1)
	check([]int{0, 0, 0, 0}, -1)

	check([]int{1, 0}, 1)
	check([]int{0, 1}, 2)
	check([]int{1, 1}, 3)

	_Dimension = 3
	check([]int{0}, -1)
	check([]int{0, 0}, -1)
	check([]int{0, 0, 0}, 0)
	check([]int{0, 0, 0, 0}, -1)

	check([]int{1, 0, 0}, 1)
	check([]int{0, 1, 0}, 2)
	check([]int{0, 0, 1}, 4)

	check([]int{1, 1, 1}, 7)
	check([]int{0, 1, 1}, 6)

	check([]int{5, 0, 0}, -2)
	check([]int{0, 5, 0}, -2)
	check([]int{0, 0, 5}, -2)

	_Size = 3
	check([]int{5, 0, 0}, -2)
	check([]int{0, 5, 0}, -2)
	check([]int{0, 0, 5}, -2)

	check([]int{1, 0, 0}, 1)
	check([]int{0, 1, 0}, 3)
	check([]int{0, 0, 1}, 9)

	check([]int{0, 2, 2}, 24)
	check([]int{2, 0, 3}, -2)
	check([]int{1, 2, 1}, 16)
}

func TestCubeParsePosition(t *testing.T) {
	check := func(coord int, want []int) {
		have := parsePosition(coord)

		if !reflect.DeepEqual(have, want) {
			t.Errorf("WANT %v HAVE %v", want, have)
		}
	}

	_Size = 2
	_Dimension = 2
	check(-1, []int{})
	check(int(math.Pow(float64(_Size), float64(_Dimension)))+1, []int{})
	check(0, []int{0, 0})
	check(1, []int{1, 0})
	check(2, []int{0, 1})
	check(3, []int{1, 1})
	check(4, []int{})

	_Size = 3
	_Dimension = 3

	check(8, []int{2, 2, 0})
	check(26, []int{2, 2, 2})
	check(27, []int{})
}

func TestFindIndexes(t *testing.T) {
	check := func(p _Plane, want1, want2 int) {
		have1, have2 := findIndexes(p)
		if have1 != want1 || have2 != want2 {
			t.Errorf("WANT %v %v HAVE %v %v\tfor %v", want1, want2, have1, have2, p)
		}
	}

	check(_Plane{directionX, directionY}, 0, 1)
	check(_Plane{directionX, directionNY}, 0, 1)
	check(_Plane{directionX, directionZ}, 0, 2)
	check(_Plane{directionX, directionNZ}, 0, 2)
	check(_Plane{directionNY, directionNZ}, 1, 2)
	check(_Plane{directionZ, directionNX}, 2, 0)
}

// func TestMove(t *testing.T) {
// 	check := func(start []Box, position []int, direction _Plane, want []Box) {
// 		c := Cube{}
// 		copy(c.arr, start)

// 		c.Move(position, direction)
// 		// if c.arr != want {

// 		// }
// 	}

// 	size = 3
// 	dimension = 3
// 	check()
// }

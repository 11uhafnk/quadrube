package main

import (
	"testing"
)

func TestCubeGetPosition(t *testing.T) {

	check := func(coords []int, want int) {
		have := getPosition(coords...)
		if have != want {
			t.Errorf("WANT %v HAVE %v", have, want)
		}
	}

	dimension = 2
	check([]int{0}, -1)
	check([]int{0, 0}, 0)
	check([]int{0, 0, 0}, -1)
	check([]int{0, 0, 0, 0}, -1)

	check([]int{1, 0}, 1)
	check([]int{0, 1}, 2)
	check([]int{1, 1}, 3)

	dimension = 3
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

	size = 3
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

package main

import (
	"errors"
	"testing"
)

func BenchmarkCheck1(b *testing.B) {
	d := directionX

	for ii := 0; ii < b.N; ii++ {
		d.Check()
	}
}

func TestDirectionReverse(t *testing.T) {

	check := func(
		check _Direction,
		want _Direction,
	) {
		check = check.Reverse()
		if check != want {
			t.Errorf("WANT %v HAVE %v", want, check)
		}
	}

	check(directionX, directionNX)
	check(directionNX, directionX)
	check(directionY, directionNY)
	check(directionZ, directionNZ)
	check(directionV, directionV) // because it large than dimention

	check(directionX|directionY, directionNX|directionNY)
}

func TestDirectionCheck(t *testing.T) {

	check := func(d _Direction, want error) {
		err := d.Check()
		if err != want && !errors.Is(err, want) {
			t.Errorf("WANT %v HAVE %v", want, err)
		}
	}

	check(directionX&directionY, ErrWrongDimention)
	check(directionX|directionY, ErrWrongDimention)
	check(directionV, ErrWrongDimention)
	check(directionX, nil)
	check(directionY, nil)
	check(directionZ, nil)
}

func TestPlaneCheck(t *testing.T) {

	check := func(p _Plane, want error) {
		err := p.Check()
		if err != want && !errors.Is(err, want) {
			t.Errorf("WANT %v HAVE %v", want, err)
		}
	}

	check(_Plane{directionX, directionX}, ErrWrongDirection)
	check(_Plane{directionX, directionNX}, ErrWrongDirection)
	check(_Plane{directionX, directionV}, ErrWrongDirection)
	check(_Plane{directionX, directionX | directionY}, ErrWrongDirection)
	check(_Plane{directionX, directionX & directionY}, ErrWrongDirection)
	check(_Plane{directionX, directionY}, nil)
	check(_Plane{directionY, directionX}, nil)
	check(_Plane{directionX, directionZ}, nil)
}

func TestGetTouchSides(t *testing.T) {

	check := func(coords []int, wantTouchs int, wantDirs _Direction) {
		haveTouchs, haveDirs := getTouchSides(coords)
		if haveDirs != wantDirs {
			t.Errorf("DIRS WANT %08b HAVE %08b, from %v", wantDirs, haveDirs, coords)
		}
		if haveTouchs != wantTouchs {
			t.Errorf("TOUCHS WANT %d HAVE %d, from %v", wantTouchs, haveTouchs, coords)
		}
	}

	size = 3
	dimension = 3
	check([]int{1, 1, 1}, 0, 0)

	check([]int{0, 1, 1}, 1, directionNX)
	check([]int{2, 1, 1}, 1, directionX)
	check([]int{1, 0, 1}, 1, directionNY)
	check([]int{1, 2, 1}, 1, directionY)
	check([]int{1, 1, 0}, 1, directionNZ)
	check([]int{1, 1, 2}, 1, directionZ)

	check([]int{1, 0, 0}, 2, directionNY|directionNZ)
	check([]int{0, 0, 0}, 3, directionNX|directionNY|directionNZ)

	dimension = 4
	check([]int{1, 1, 1, 1}, 0, 0)
	check([]int{1, 1, 1, 0}, 1, directionNV)
	check([]int{1, 1, 1, 2}, 1, directionV)
	check([]int{0, 0, 0, 0}, 4, directionNX|directionNY|directionNZ|directionNV)

	dimension = 3
}

func TestReDirection(t *testing.T) {

	p := _Plane{directionX, directionY}

	check := func(coords []int, want _Plane) {
		have := p.ReDirection(coords)
		// if err != wantErr {
		// 	t.Errorf("WANT %v HAVE %v\tfor%v\n", wantErr, err, coords)
		// }
		if want != have {
			t.Errorf("WANT %+v HAVE %+vfor%v\n", want, have, coords)
		}
	}

	size = 3
	dimension = 4
	check([]int{1, 1, 0, 0}, _Plane{})

	dimension = 3

	check([]int{1, 1, 1}, _Plane{directionX, directionY})

	check([]int{1, 1, 0}, _Plane{directionX, directionY})
	check([]int{1, 1, 2}, _Plane{directionX, directionY})
	check([]int{1, 0, 1}, _Plane{directionX, directionY})
	check([]int{0, 1, 1}, _Plane{directionNY, directionX})
	check([]int{2, 1, 1}, _Plane{directionY, directionNX})
	check([]int{1, 2, 1}, _Plane{directionNX, directionNY})

	check([]int{0, 0, 1}, _Plane{directionX, directionY})
	check([]int{1, 0, 0}, _Plane{directionX, directionY})
	check([]int{1, 0, 2}, _Plane{directionX, directionY})
	check([]int{2, 2, 1}, _Plane{directionNX, directionNY})
	check([]int{1, 2, 2}, _Plane{directionNX, directionNY})
	check([]int{1, 2, 0}, _Plane{directionNX, directionNY})
	check([]int{2, 1, 2}, _Plane{directionY, directionNX})
	check([]int{2, 0, 1}, _Plane{directionY, directionNX})
	check([]int{2, 1, 0}, _Plane{directionY, directionNX})
	check([]int{0, 1, 0}, _Plane{directionNY, directionX})
	check([]int{0, 2, 1}, _Plane{directionNY, directionX})
	check([]int{0, 1, 2}, _Plane{directionNY, directionX})

	check([]int{0, 0, 0}, _Plane{directionX, directionY})
	check([]int{0, 0, 2}, _Plane{directionX, directionY})
	check([]int{2, 0, 0}, _Plane{directionY, directionNX})
	check([]int{2, 0, 2}, _Plane{directionY, directionNX})
	check([]int{0, 2, 0}, _Plane{directionNY, directionX})
	check([]int{0, 2, 2}, _Plane{directionNY, directionX})
	check([]int{2, 2, 0}, _Plane{directionNX, directionNY})
	check([]int{2, 2, 2}, _Plane{directionNX, directionNY})
}

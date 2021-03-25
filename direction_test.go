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

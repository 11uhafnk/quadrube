//
package main

import (
	"errors"
	"testing"
)

func TestSideMove(t *testing.T) {
	t.Parallel()

	testMove := func(t *testing.T, b Box, to _Plane, b2 Box, err error) {
		from := b.Get()
		err2 := b.Move(to)
		if !errors.Is(err2, err) {
			t.Errorf("%v to %v ) WANT ERR: %v GOT %v", from, to, err, err2)
		}
		// if err != nil {
		// 	return
		// }
		if b.Get() != b2.Get() {
			t.Errorf("%v to %v ) WANT %v GOT %v", from, to, b2.Get(), b.Get())
		}
	}

	testMove(t,
		&Side{Color{color: ColorRed, orientation: directionNX}},
		_Plane{directionNX, directionNX},
		&Side{Color{color: ColorRed, orientation: directionNX}},
		ErrWrongDirection,
	)
	testMove(t,
		&Side{Color{color: ColorRed, orientation: directionNX}},
		_Plane{directionNX, directionX},
		&Side{Color{color: ColorRed, orientation: directionNX}},
		ErrWrongDirection,
	)
	testMove(t,
		&Side{Color{color: ColorRed, orientation: directionNX}},
		_Plane{directionNX, directionV},
		&Side{Color{color: ColorRed, orientation: directionNX}},
		ErrWrongDirection,
	)

	// X
	testMove(t,
		&Side{Color{color: ColorRed, orientation: directionNX}},
		_Plane{directionX, directionY},
		&Side{Color{color: ColorRed, orientation: directionNX}},
		ErrMove,
	)
	testMove(t,
		&Side{Color{color: ColorRed, orientation: directionNX}},
		_Plane{directionX, directionNY},
		&Side{Color{color: ColorRed, orientation: directionNX}},
		ErrMove,
	)
	testMove(t,
		&Side{Color{color: ColorRed, orientation: directionNX}},
		_Plane{directionX, directionZ},
		&Side{Color{color: ColorRed, orientation: directionNX}},
		ErrMove,
	)
	testMove(t,
		&Side{Color{color: ColorRed, orientation: directionNX}},
		_Plane{directionX, directionNZ},
		&Side{Color{color: ColorRed, orientation: directionNX}},
		ErrMove,
	)
	// NX
	testMove(t,
		&Side{Color{color: ColorRed, orientation: directionNX}},
		_Plane{directionNX, directionY},
		&Side{Color{color: ColorRed, orientation: directionNX}},
		ErrMove,
	)
	testMove(t,
		&Side{Color{color: ColorRed, orientation: directionNX}},
		_Plane{directionNX, directionNY},
		&Side{Color{color: ColorRed, orientation: directionNX}},
		ErrMove,
	)
	testMove(t,
		&Side{Color{color: ColorRed, orientation: directionNX}},
		_Plane{directionNX, directionZ},
		&Side{Color{color: ColorRed, orientation: directionNX}},
		ErrMove,
	)
	testMove(t,
		&Side{Color{color: ColorRed, orientation: directionNX}},
		_Plane{directionNX, directionNZ},
		&Side{Color{color: ColorRed, orientation: directionNX}},
		ErrMove,
	)

	// Y
	testMove(t,
		&Side{Color{color: ColorRed, orientation: directionNX}},
		_Plane{directionY, directionX},
		&Side{Color{color: ColorRed, orientation: directionY}},
		nil,
	)
	testMove(t,
		&Side{Color{color: ColorRed, orientation: directionNX}},
		_Plane{directionY, directionNX},
		&Side{Color{color: ColorRed, orientation: directionNX}},
		ErrMove,
	)
	testMove(t,
		&Side{Color{color: ColorRed, orientation: directionNX}},
		_Plane{directionY, directionZ},
		&Side{Color{color: ColorRed, orientation: directionNX}},
		nil,
	)
	testMove(t,
		&Side{Color{color: ColorRed, orientation: directionNX}},
		_Plane{directionY, directionNZ},
		&Side{Color{color: ColorRed, orientation: directionNX}},
		nil,
	)
	// NY
	testMove(t,
		&Side{Color{color: ColorRed, orientation: directionNX}},
		_Plane{directionNY, directionX},
		&Side{Color{color: ColorRed, orientation: directionNY}},
		nil,
	)
	testMove(t,
		&Side{Color{color: ColorRed, orientation: directionNX}},
		_Plane{directionNY, directionNX},
		&Side{Color{color: ColorRed, orientation: directionNX}},
		ErrMove,
	)
	testMove(t,
		&Side{Color{color: ColorRed, orientation: directionNX}},
		_Plane{directionNY, directionZ},
		&Side{Color{color: ColorRed, orientation: directionNX}},
		nil,
	)
	testMove(t,
		&Side{Color{color: ColorRed, orientation: directionNX}},
		_Plane{directionNY, directionNZ},
		&Side{Color{color: ColorRed, orientation: directionNX}},
		nil,
	)

	// Y
	testMove(t,
		&Side{Color{color: ColorRed, orientation: directionNX}},
		_Plane{directionZ, directionX},
		&Side{Color{color: ColorRed, orientation: directionZ}},
		nil,
	)
	testMove(t,
		&Side{Color{color: ColorRed, orientation: directionNX}},
		_Plane{directionZ, directionNX},
		&Side{Color{color: ColorRed, orientation: directionNX}},
		ErrMove,
	)
	testMove(t,
		&Side{Color{color: ColorRed, orientation: directionNX}},
		_Plane{directionZ, directionY},
		&Side{Color{color: ColorRed, orientation: directionNX}},
		nil,
	)
	testMove(t,
		&Side{Color{color: ColorRed, orientation: directionNX}},
		_Plane{directionZ, directionNY},
		&Side{Color{color: ColorRed, orientation: directionNX}},
		nil,
	)
	// NY
	testMove(t,
		&Side{Color{color: ColorRed, orientation: directionNX}},
		_Plane{directionNZ, directionX},
		&Side{Color{color: ColorRed, orientation: directionNZ}},
		nil,
	)
	testMove(t,
		&Side{Color{color: ColorRed, orientation: directionNX}},
		_Plane{directionNZ, directionNX},
		&Side{Color{color: ColorRed, orientation: directionNX}},
		ErrMove,
	)
	testMove(t,
		&Side{Color{color: ColorRed, orientation: directionNX}},
		_Plane{directionNZ, directionY},
		&Side{Color{color: ColorRed, orientation: directionNX}},
		nil,
	)
	testMove(t,
		&Side{Color{color: ColorRed, orientation: directionNX}},
		_Plane{directionNZ, directionNY},
		&Side{Color{color: ColorRed, orientation: directionNX}},
		nil,
	)
}

func TestEdgeMove(t *testing.T) {
	t.Parallel()

	testMove := func(t *testing.T, b Box, to _Plane, b2 Box, err error) {
		from := b.Get()
		err2 := b.Move(to)
		if !errors.Is(err2, err) {
			t.Errorf("%v to %v ) WANT ERR: %v GOT %v", from, to, err, err2)
		}
		// if err != nil {
		// 	return
		// }
		if b.Get() != b2.Get() {
			t.Errorf("%v to %v ) WANT %v GOT %v", from, to, b2.Get(), b.Get())
		}
	}

	testMove(t,
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		_Plane{directionNX, directionNX},
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		ErrWrongDirection,
	)
	testMove(t,
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		_Plane{directionNX, directionX},
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		ErrWrongDirection,
	)
	testMove(t,
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		_Plane{directionNX, directionV},
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		ErrWrongDirection,
	)

	func() {
		defer func() {
			if err := recover(); err == nil {
				t.Errorf("haven't panic")
			}
		}()
		testMove(t,
			&Edge{[]Color{{ColorRed, directionNX}}},
			_Plane{directionNX, directionY},
			&Edge{[]Color{{ColorRed, directionNX}}},
			ErrWrongDirection,
		)
	}()

	// X
	testMove(t,
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		_Plane{directionX, directionY},
		&Edge{[]Color{{ColorRed, directionNY}, {ColorGreen, directionX}}},
		nil,
	)
	testMove(t,
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		_Plane{directionX, directionNY},
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		ErrMove,
	)
	testMove(t,
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		_Plane{directionX, directionZ},
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		ErrMove,
	)
	testMove(t,
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		_Plane{directionX, directionNZ},
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		ErrMove,
	)
	// NX
	testMove(t,
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		_Plane{directionNX, directionY},
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		ErrMove,
	)
	testMove(t,
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		_Plane{directionNX, directionNY},
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		ErrMove,
	)
	testMove(t,
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		_Plane{directionNX, directionZ},
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		ErrMove,
	)
	testMove(t,
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		_Plane{directionNX, directionNZ},
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		ErrMove,
	)

	// Y
	testMove(t,
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		_Plane{directionY, directionX},
		&Edge{[]Color{{ColorRed, directionY}, {ColorGreen, directionNX}}},
		nil,
	)
	testMove(t,
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		_Plane{directionY, directionNX},
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		ErrMove,
	)
	testMove(t,
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		_Plane{directionY, directionZ},
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		ErrMove,
	)
	testMove(t,
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		_Plane{directionY, directionNZ},
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		ErrMove,
	)
	// NY
	testMove(t,
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		_Plane{directionNY, directionX},
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		ErrMove,
	)
	testMove(t,
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		_Plane{directionNY, directionNX},
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		ErrMove,
	)
	testMove(t,
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		_Plane{directionNY, directionZ},
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		ErrMove,
	)
	testMove(t,
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		_Plane{directionNY, directionNZ},
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		ErrMove,
	)

	// Z
	testMove(t,
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		_Plane{directionZ, directionX},
		&Edge{[]Color{{ColorRed, directionZ}, {ColorGreen, directionNY}}},
		nil,
	)
	testMove(t,
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		_Plane{directionZ, directionNX},
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		ErrMove,
	)
	testMove(t,
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		_Plane{directionZ, directionY},
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionZ}}},
		nil,
	)
	testMove(t,
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		_Plane{directionZ, directionNY},
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		ErrMove,
	)
	// NZ
	testMove(t,
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		_Plane{directionNZ, directionX},
		&Edge{[]Color{{ColorRed, directionNZ}, {ColorGreen, directionNY}}},
		nil,
	)
	testMove(t,
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		_Plane{directionNZ, directionNX},
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		ErrMove,
	)
	testMove(t,
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		_Plane{directionNZ, directionY},
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNZ}}},
		nil,
	)
	testMove(t,
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		_Plane{directionNZ, directionNY},
		&Edge{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}}},
		ErrMove,
	)
}

func TestVertexMove(t *testing.T) {
	t.Parallel()

	testMove := func(t *testing.T, b Box, to _Plane, b2 Box, err error) {
		from := b.Get()
		err2 := b.Move(to)
		if !errors.Is(err2, err) {
			t.Errorf("%v to %v ) WANT ERR: %v GOT %v", from, to, err, err2)
		}
		// if err != nil {
		// 	return
		// }
		if b.Get() != b2.Get() {
			t.Errorf("%v to %v ) WANT %v GOT %v", from, to, b2.Get(), b.Get())
		}
	}

	testMove(t,
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		_Plane{directionNX, directionNX},
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		ErrWrongDirection,
	)
	testMove(t,
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		_Plane{directionNX, directionX},
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		ErrWrongDirection,
	)
	testMove(t,
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		_Plane{directionNX, directionV},
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		ErrWrongDirection,
	)

	func() {
		defer func() {
			if err := recover(); err == nil {
				t.Errorf("haven't panic")
			}
		}()
		testMove(t,
			&Vertex{[]Color{{ColorRed, directionNX}}},
			_Plane{directionNX, directionY},
			&Vertex{[]Color{{ColorRed, directionNX}}},
			ErrWrongDirection,
		)
	}()

	// X
	testMove(t,
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		_Plane{directionX, directionY},
		&Vertex{[]Color{{ColorRed, directionNY}, {ColorGreen, directionX}, {ColorBlue, directionNZ}}},
		nil,
	)
	testMove(t,
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		_Plane{directionX, directionNY},
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		ErrMove,
	)
	testMove(t,
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		_Plane{directionX, directionZ},
		&Vertex{[]Color{{ColorRed, directionNZ}, {ColorGreen, directionNY}, {ColorBlue, directionX}}},
		nil,
	)
	testMove(t,
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		_Plane{directionX, directionNZ},
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		ErrMove,
	)
	// NX
	testMove(t,
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		_Plane{directionNX, directionY},
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		ErrMove,
	)
	testMove(t,
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		_Plane{directionNX, directionNY},
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		ErrMove,
	)
	testMove(t,
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		_Plane{directionNX, directionZ},
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		ErrMove,
	)
	testMove(t,
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		_Plane{directionNX, directionNZ},
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		ErrMove,
	)

	// Y
	testMove(t,
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		_Plane{directionY, directionX},
		&Vertex{[]Color{{ColorRed, directionY}, {ColorGreen, directionNX}, {ColorBlue, directionNZ}}},
		nil,
	)
	testMove(t,
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		_Plane{directionY, directionNX},
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		ErrMove,
	)
	testMove(t,
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		_Plane{directionY, directionZ},
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNZ}, {ColorBlue, directionY}}},
		nil,
	)
	testMove(t,
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		_Plane{directionY, directionNZ},
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		ErrMove,
	)
	// NY
	testMove(t,
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		_Plane{directionNY, directionX},
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		ErrMove,
	)
	testMove(t,
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		_Plane{directionNY, directionNX},
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		ErrMove,
	)
	testMove(t,
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		_Plane{directionNY, directionZ},
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		ErrMove,
	)
	testMove(t,
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		_Plane{directionNY, directionNZ},
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		ErrMove,
	)

	// Z
	testMove(t,
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		_Plane{directionZ, directionX},
		&Vertex{[]Color{{ColorRed, directionZ}, {ColorGreen, directionNY}, {ColorBlue, directionNX}}},
		nil,
	)
	testMove(t,
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		_Plane{directionZ, directionNX},
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		ErrMove,
	)
	testMove(t,
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		_Plane{directionZ, directionY},
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionZ}, {ColorBlue, directionNY}}},
		nil,
	)
	testMove(t,
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		_Plane{directionZ, directionNY},
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		ErrMove,
	)
	// NZ
	testMove(t,
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		_Plane{directionNZ, directionX},
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		ErrMove,
	)
	testMove(t,
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		_Plane{directionNZ, directionNX},
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		ErrMove,
	)
	testMove(t,
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		_Plane{directionNZ, directionY},
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		ErrMove,
	)
	testMove(t,
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		_Plane{directionNZ, directionNY},
		&Vertex{[]Color{{ColorRed, directionNX}, {ColorGreen, directionNY}, {ColorBlue, directionNZ}}},
		ErrMove,
	)
}

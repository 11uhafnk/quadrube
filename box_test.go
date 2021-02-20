package main

import (
	"testing"
)

// pair ...
type pair struct {
	to     _Plane
	result Box
	err    error
}

func TestSideMove(t *testing.T) {

	b := &Side{
		colors: Color{ColorRed, directionNX},
	}

	testMove(t, b, 1,
		_Plane{directionNX, directionNY},
		&Side{Color{color: ColorRed, orientation: directionNX}},
		nil,
	)

}

func testMove(t *testing.T, b Box, ii int, to _Plane, b2 Box, err error) {
	err2 := b.Move(to)
	if err2 != err {
		t.Errorf("%02d) WANT ERR: %v GOT %v", ii, err, err2)
	}
	if err != nil {
		return
	}
	if b.Get() != b2.Get() {
		t.Errorf("%02d) WANT %v GOT %v", ii, b2.Get(), b.Get())
	}
}

// 	fmt.Printf("\nWrong:\n")
// 	b.Move([2]_Direction{directionX, directionX})
// 	// e := ErrMove
// 	// if errors.As(err, &e) {
// 	// 	panicOnError(err)
// 	// }
// 	b = newb()
// 	b.Move([2]_Direction{directionNX, directionX})
// 	fmt.Printf("\n")

// 	fmt.Printf("\nALL:\n")
// 	fmt.Printf("x:\n")
// 	b = newb()
// 	b.Move([2]_Direction{directionX, directionY})
// 	b = newb()
// 	b.Move([2]_Direction{directionX, directionNY})
// 	b = newb()
// 	b.Move([2]_Direction{directionX, directionZ})
// 	b = newb()
// 	b.Move([2]_Direction{directionX, directionNZ})

// 	fmt.Printf("not x:\n")
// 	b = newb()
// 	b.Move([2]_Direction{directionNX, directionY})
// 	b = newb()
// 	b.Move([2]_Direction{directionNX, directionNY})
// 	b = newb()
// 	b.Move([2]_Direction{directionNX, directionZ})
// 	b = newb()
// 	b.Move([2]_Direction{directionNX, directionNZ})

// 	fmt.Printf("y:\n")
// 	b = newb()
// 	b.Move([2]_Direction{directionY, directionX})
// 	b = newb()
// 	b.Move([2]_Direction{directionY, directionNX})
// 	b = newb()
// 	b.Move([2]_Direction{directionY, directionZ})
// 	b = newb()
// 	b.Move([2]_Direction{directionY, directionNZ})

// 	fmt.Printf("not y:\n")
// 	b = newb()
// 	b.Move([2]_Direction{directionNY, directionX})
// 	b = newb()
// 	b.Move([2]_Direction{directionNY, directionNX})
// 	b = newb()
// 	b.Move([2]_Direction{directionNY, directionZ})
// 	b = newb()
// 	b.Move([2]_Direction{directionNY, directionNZ})

// 	fmt.Printf("z:\n")
// 	b = newb()
// 	b.Move([2]_Direction{directionZ, directionX})
// 	b = newb()
// 	b.Move([2]_Direction{directionZ, directionNX})
// 	b = newb()
// 	b.Move([2]_Direction{directionZ, directionY})
// 	b = newb()
// 	b.Move([2]_Direction{directionZ, directionNY})

// 	fmt.Printf("not z:\n")
// 	b = newb()
// 	b.Move([2]_Direction{directionNZ, directionX})
// 	b = newb()
// 	b.Move([2]_Direction{directionNZ, directionNX})
// 	b = newb()
// 	b.Move([2]_Direction{directionNZ, directionY})
// 	b = newb()
// 	b.Move([2]_Direction{directionNZ, directionNY})

// }

package main

import (
	"fmt"
	"testing"
)

// pair ...
type pair struct {
	dir    _Direction
	result _Direction
}

func TestSideMove(t *testing.T) {

	pairs := {}
}

func testMove(t *testing.T, b box, directions ...pair) {
	newb := func() box {
		// return &Side{
		// 	colors: Color{ColorRed, directionNZ},
		// }
		return &Vertex{
			// return &Edge{
			colors: []Color{
				{ColorGreen, directionNX},
				{ColorRed, directionNZ},
				{ColorBlue, directionNY},
				// {ColorCyan, directionNV},
			},
		}
	}

	// p := _Plane{directionY, directionNX | directionX}

	// fmt.Println(p[0].Check(), p[0])
	// fmt.Println(p[1].Check(), p[1])
	// fmt.Println(p.Check(), p)
	// return

	b := newb()
	// var err error
	fmt.Println(b.Get())

	fmt.Printf("\nWrong:\n")
	b.Move([2]_Direction{directionX, directionX})
	// e := ErrMove
	// if errors.As(err, &e) {
	// 	panicOnError(err)
	// }
	b = newb()
	b.Move([2]_Direction{directionNX, directionX})
	fmt.Printf("\n")

	fmt.Printf("\nALL:\n")
	fmt.Printf("x:\n")
	b = newb()
	b.Move([2]_Direction{directionX, directionY})
	b = newb()
	b.Move([2]_Direction{directionX, directionNY})
	b = newb()
	b.Move([2]_Direction{directionX, directionZ})
	b = newb()
	b.Move([2]_Direction{directionX, directionNZ})

	fmt.Printf("not x:\n")
	b = newb()
	b.Move([2]_Direction{directionNX, directionY})
	b = newb()
	b.Move([2]_Direction{directionNX, directionNY})
	b = newb()
	b.Move([2]_Direction{directionNX, directionZ})
	b = newb()
	b.Move([2]_Direction{directionNX, directionNZ})

	fmt.Printf("y:\n")
	b = newb()
	b.Move([2]_Direction{directionY, directionX})
	b = newb()
	b.Move([2]_Direction{directionY, directionNX})
	b = newb()
	b.Move([2]_Direction{directionY, directionZ})
	b = newb()
	b.Move([2]_Direction{directionY, directionNZ})

	fmt.Printf("not y:\n")
	b = newb()
	b.Move([2]_Direction{directionNY, directionX})
	b = newb()
	b.Move([2]_Direction{directionNY, directionNX})
	b = newb()
	b.Move([2]_Direction{directionNY, directionZ})
	b = newb()
	b.Move([2]_Direction{directionNY, directionNZ})

	fmt.Printf("z:\n")
	b = newb()
	b.Move([2]_Direction{directionZ, directionX})
	b = newb()
	b.Move([2]_Direction{directionZ, directionNX})
	b = newb()
	b.Move([2]_Direction{directionZ, directionY})
	b = newb()
	b.Move([2]_Direction{directionZ, directionNY})

	fmt.Printf("not z:\n")
	b = newb()
	b.Move([2]_Direction{directionNZ, directionX})
	b = newb()
	b.Move([2]_Direction{directionNZ, directionNX})
	b = newb()
	b.Move([2]_Direction{directionNZ, directionY})
	b = newb()
	b.Move([2]_Direction{directionNZ, directionNY})

}

// Equal description
func equalBox(a, b box) bool {
	return true
}

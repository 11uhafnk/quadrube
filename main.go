package main

import "fmt"

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	newb := func() box {
		return &Edge{
			colors: []Color{
				{ColorGreen, directionNX},
				{ColorRed, directionNZ},
			},
		}
	}

	// p := _Plane{directionY, directionNX | directionX}

	// fmt.Println(p[0].Check(), p[0])
	// fmt.Println(p[1].Check(), p[1])
	// fmt.Println(p.Check(), p)
	// return

	b := newb()
	fmt.Println(b.Get())

	fmt.Printf("\nWrong:")
	b.Move([2]_Direction{directionX, directionX})
	b = newb()
	b.Move([2]_Direction{directionNX, directionX})
	fmt.Printf("\n")

	fmt.Printf("\nInner:")
	b = newb()
	b.Move([2]_Direction{directionX, directionZ})
	b = newb()
	b.Move([2]_Direction{directionX, directionNZ})
	b = newb()
	b.Move([2]_Direction{directionZ, directionX})

	fmt.Printf("\nOuter:")
	b = newb()
	b.Move([2]_Direction{directionY, directionZ})
	b = newb()
	b.Move([2]_Direction{directionNY, directionZ})
	b = newb()
	b.Move([2]_Direction{directionY, directionX})
	b = newb()
	b.Move([2]_Direction{directionNY, directionX})

	fmt.Printf("\n")

	fmt.Println("end")
}

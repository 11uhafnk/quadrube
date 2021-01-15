package main

import "fmt"

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	b := Side{
		colors: Color{ColorGreen, directionNY},
	}

	d := directionX | directionNY | directionV

	fmt.Println(d.Check())
	return

	fmt.Println(b.Get())

	b.Move(directionZ | directionNX)

	b.Move(directionX)
	b.Move(directionY)
	b.Move(directionNX)
	b.Move(directionZ)
	b.Move(directionNY)
	b.Move(directionNZ)
	b.Move(directionV)
	b.Move(directionNV)

	fmt.Println(b.Get())

	fmt.Printf("\n")

	fmt.Println("end")
}

package main

import (
	"fmt"
)

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error

	c := InitCube(3, 3)

	c.selected = 13
	c.Print()
	c.selected -= _Size
	c.Print()
	c.selected -= _Size
	c.Print()
	// fmt.Println("\n")

	//	err = c.Move([]int{0, 0, 1}, _Plane{directionY, directionX})
	//	panicOnError(err)
	//	c.Print()

	//	err = c.Move([]int{0, 0, 1}, _Plane{directionX, directionY})
	//	panicOnError(err)
	//	c.Print()

	err = c.Move([]int{0, 0, 0}, _Plane{directionX, directionZ})
	panicOnError(err)
	c.Print()

	fmt.Println("end")
}

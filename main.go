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

	c := InitCube(3, 4)

	c.Print()
	// fmt.Println("\n")

	err := c.Move([]int{0, 0, 1}, _Plane{directionX, directionY})
	panicOnError(err)

	c.Print()

	fmt.Println("end")
}

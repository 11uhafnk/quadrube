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

	c3 := InitCube(3, 3)
	c3.Print()
	fmt.Println("\n---------------------------")

	// c.selected = 13
	// c.Print()
	// c.selected -= _Size
	// c.Print()
	// c.selected -= _Size
	// c.Print()
	// // fmt.Println("\n")

	err = c3.Move([]int{0, 0, 1}, _Plane{directionY, directionX})
	panicOnError(err)
	c3.Print()
	fmt.Println("\n---------------------------")

	err = c3.Move([]int{0, 0, 1}, _Plane{directionX, directionY})
	panicOnError(err)
	c3.Print()
	fmt.Println("\n---------------------------")

	err = c3.Move([]int{0, 0, 0}, _Plane{directionX, directionZ})
	panicOnError(err)
	c3.Print()
	fmt.Println("\n---------------------------")

	// 4d cube
	c4 := InitCube(4, 3)
	c4.Print()
	fmt.Println("\n---------------------------")

	c4.Move([]int{0, 0, 0, 0}, _Plane{directionX, directionY})
	c4.Print()
	fmt.Println("\n---------------------------")

	fmt.Println("end")
}

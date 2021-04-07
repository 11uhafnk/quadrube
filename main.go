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

	c := InitCube(5, 3)

	c.Print()

	fmt.Println("end")
}

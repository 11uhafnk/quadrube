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


	fmt.Printf("\n")

	fmt.Println("end")
}

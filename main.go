package main

import "fmt"

func main() {

	b := Side{
		colors: Color{ColorRed, 0},
	}

	fmt.Println(b.Get())

	b.Move(0, true)

	fmt.Println(b.Get())

	fmt.Println("end")

	fmt.Printf("%08b %08b %v\n", positionX, positionX&maskPositionZ, positionX&maskPositionZ != 0)
	fmt.Printf("%08b %08b %v\n", positionNX, positionNX&maskPositionZ, positionNX&maskPositionZ != 0)
	fmt.Printf("%08b %08b %v\n", positionY, positionY&maskPositionZ, positionY&maskPositionZ != 0)
	fmt.Printf("%08b %08b %v\n", positionNY, positionNY&maskPositionZ, positionNY&maskPositionZ != 0)
	fmt.Printf("%08b %08b %v\n", positionZ, positionZ&maskPositionZ, positionZ&maskPositionZ != 0)
	fmt.Printf("%08b %08b %v\n", positionNZ, positionNZ&maskPositionZ, positionNZ&maskPositionZ != 0)
	fmt.Printf("%08b %08b %v\n", positionV, positionV&maskPositionZ, positionV&maskPositionZ != 0)
	fmt.Printf("%08b %08b %v\n", positionNV, positionNV&maskPositionZ, positionNV&maskPositionZ != 0)
	fmt.Printf("\n")

	fmt.Printf("%08b %v\n", maskPositionX, maskPositionX)
	fmt.Printf("%08b %v\n", maskPositionY, maskPositionY)
	fmt.Printf("%08b %v\n", maskPositionZ, maskPositionZ)
	fmt.Printf("%08b %v\n", maskPositionV, maskPositionV)
	fmt.Printf("\n")

	fmt.Println("end")
}

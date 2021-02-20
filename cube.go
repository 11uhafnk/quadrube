package main

import "math"

// Cube ...
type Cube struct {
	// it as [][][][]box
	//        ^ deep as dimention
	arr []Box
}

// Rotate description
func (c *Cube) Rotate(
	position []int,
	direction _Plane,
) {
	return
}

// getPosition( X Y Z V ...)
func getPosition(
	coords ...int,
) (

	coord int,
) {
	if len(coords) != dimension {
		return -1
	}

	for ii := 0; ii < dimension; ii++ {
		coord += int(math.Pow(float64(size), float64(ii))) * coords[ii]
	}

	return coord
}

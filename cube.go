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
		if coords[ii] >= dimension {
			return -2
		}
		coord += int(math.Pow(float64(size), float64(ii))) * coords[ii]
	}

	if coord >= int(math.Pow(float64(size), float64(dimension))) {
		return -3
	}

	return coord
}

package main

// 3    1    2
//  V   Y   Z
//   ^  ^  ^
//    \ | /
//     \|/
//      ------> X - 0

var (
	dimension = 3

	dimensionMask = 0

	size = 2
)

func init() {
	for ii := 0; ii < dimension; ii++ {
		dimensionMask |= int(directions[ii])
	}
}

type _Direction int

const (
	directionX _Direction = 1 << iota
	directionNX
	directionY
	directionNY
	directionZ
	directionNZ
	directionV
	directionNV
)

const (
	maskDirectionX = 3 << (2 * iota)
	maskDirectionY
	maskDirectionZ
	maskDirectionV
)

var directions = []_Direction{maskDirectionX, maskDirectionY, maskDirectionZ, maskDirectionV}

const directionLetters = "XYZV"

func (d _Direction) String() string {
	if d == 0 {
		return "none"
	}
	result := ""
	for ii := 0; ii < dimension; ii++ {
		if d&maskDirectionX != 0 {
			result += directionLetters[ii : ii+1]
			if d&directionX != 0 {
				result += "+"
			} else {
				result += "-"
			}
		}

		d = d >> 2
	}

	return result
}

func (d _Direction) Reverse() _Direction {
	for ii := 0; ii < dimension; ii++ {
		if d&directions[ii] != 0 {
			d ^= directions[ii]
		}
	}

	return d
}

func (d _Direction) Check() bool {

	d &= _Direction(dimensionMask)
	if d == 0 {
		return false
	}

	cnt := 0
	for ii := 0; ii < dimension*2; ii++ {
		if d&directionX != 0 {
			cnt++
		}
		d >>= 1
	}

	return cnt == 1
}

type _Plane [2]_Direction

func (p _Plane) Check() bool {

	if p[0]&p[1] != 0 {
		return false
	}

	xor := p[0] ^ p[1]
	for ii := 0; ii < dimension; ii++ {
		if xor == directions[ii] {
			return false
		}
	}
	return p[0].Check() && p[1].Check()
}
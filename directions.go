package main

import "fmt"

// 3    1    2
//  V   Y   Z
//   ^  ^  ^
//    \ | /
//     \|/
//      ------> X - 0

var (
	// размерность
	dimension = 3

	// маска для проверки принадлежности текущей размерности
	dimensionMask = 0

	// размер кубика в боксах
	size = 2
)

// инициализация маски размерности
func init() {
	for ii := 0; ii < dimension; ii++ {
		dimensionMask |= int(directions[ii])
	}
}

// _Direction Вектор или скорее направление
type _Direction int

// возможные направления
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

// вспомогательные маски для проверок принадлежности векторов направлений конкретные координаты
const (
	maskDirectionX = 3 << (2 * iota)
	maskDirectionY
	maskDirectionZ
	maskDirectionV
)

var directions = []_Direction{maskDirectionX, maskDirectionY, maskDirectionZ, maskDirectionV}

const directionLetters = "XYZV"

// String предварительный красивый вывод вектора для консоли
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

// Reverse меняет направление вектора на противоположное
func (d _Direction) Reverse() _Direction {
	for ii := 0; ii < dimension; ii++ {
		if d&directions[ii] != 0 {
			d ^= directions[ii]
		}
	}

	return d
}

// Проверка корректного состаяния вектора направления
func (d _Direction) Check() error {

	// вектор должен быть в текущей размерности
	d &= _Direction(dimensionMask)
	if d == 0 {
		return ErrWrongDimention
	}

	// вектор должен указывать только по одной из координат
	cnt := 0
	for ii := 0; ii < dimension*2; ii++ {
		if d&directionX != 0 {
			cnt++
		}
		d >>= 1
	}

	if cnt != 1 {
		return fmt.Errorf("direction dimention must be 1, %w", ErrWrongDimention)
	}

	return nil
}

// _Plane пока для представления плоскости вращения
// возможно потом появятся и другие назначения
type _Plane [2]_Direction

func (p _Plane) Check() error {

	// сонаправленные векторы не дают однозначной плоскости
	// оба вектора находятся в одной координате и указывают в одном направлении
	//  аля X+ X+
	if p[0]&p[1] != 0 {
		return ErrWrongDirection
	}

	// параллельные разнонаправленные векторы не дают однозначной плоскости
	// оба вектора находятся в одной координате хоть и указывают в ранзых направлениях
	//  аля X+ X-
	xor := p[0] ^ p[1]
	for ii := 0; ii < dimension; ii++ {
		if xor == directions[ii] {
			return ErrWrongDirection
		}
	}

	if err1, err2 := p[0].Check(), p[1].Check(); err1 != nil || err2 != nil {
		return fmt.Errorf("%v: %w", p, ErrWrongDirection)
	}

	return nil
}

// ReDirection возвращает направление вращения в той же плоскости
// но в удомном виде для указанной позиции
func (p _Plane) ReDirection(
	coords []int,
) _Plane {
	touchSides, dirs := getTouchSides(coords)
	_ = dirs
	switch touchSides {

	// координаты не касаются сторон
	// так что без разницы
	case 0:
		return p

	// это элемент стороны а значит
	// направление вращения должно начинаться с направления лежащего в плоскости стороны где находится данный блок
	// если блок находится на плоскости вращения ( ни один из векторов вращения
	// не совпадает со стороной касания ) то без разницы?
	case 1:

		if dirs&p[0] != 0 {
			return _Plane{p[1], p[0].Reverse()}
		} else if dirs&p[0].Reverse() != 0 {
			return _Plane{p[1].Reverse(), p[0]}
		} else if dirs&p[1] != 0 {
			return _Plane{p[0].Reverse(), p[1].Reverse()}
		} else {
			return p
		}

		// ребро
	// первый вектор вращения должне быть в плоскости касания
	case dimension - 1:

		if dirs&p[0] != 0 && dirs&p[1] == 0 {
			return _Plane{p[1], p[0].Reverse()}
		} else if dirs&p[0].Reverse() != 0 && dirs&p[1].Reverse() == 0 {
			return _Plane{p[1].Reverse(), p[0]}
		} else if dirs&p[1] != 0 && dirs&p[0].Reverse() == 0 {
			return _Plane{p[0].Reverse(), p[1].Reverse()}
		} else if dirs&p[1].Reverse() != 0 && dirs&p[0] == 0 {
			return _Plane{p[0], p[1]}
		}

	// угол
	case dimension:

		if dirs&p[0] != 0 && dirs&p[1].Reverse() != 0 {
			return _Plane{p[1], p[0].Reverse()}
		} else if dirs&p[0].Reverse() != 0 && dirs&p[1] != 0 {
			return _Plane{p[1].Reverse(), p[0]}
		} else if dirs&p[0] != 0 && dirs&p[1] != 0 {
			return _Plane{p[0].Reverse(), p[1].Reverse()}
		} else if dirs&p[0].Reverse() != 0 && dirs&p[1].Reverse() != 0 {
			return _Plane{p[0], p[1]}
		}
	}

	return _Plane{}
}

func getTouchSides(
	coords []int,
) (
	touchSides int,
	dirs _Direction,
) {
	minPos := 0
	maxPos := size - 1

	// dirs = make([]_Direction, dimension)

	for ii := range coords {
		if coords[ii] == minPos {
			dirs |= directionNX << (ii * 2)
			touchSides++
		} else if coords[ii] == maxPos {
			dirs |= directionX << (ii * 2)
			touchSides++
		}
	}
	return
}
